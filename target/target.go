package target

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"pingmaster/helpers"
	"pingmaster/user"
)

const (
	TargetType_Website = "website"
)

// Target is an interface containing each target related methods
type Target interface {

	// GetGeneric returns GenericTarget
	GetGeneric() *GenericTarget

	// GetPoolKey returns the pool key in the described format
	GetPoolKey() string

	// Ping will call the target using appropriate method
	// and return the details and error
	Ping(ctx context.Context)

	// GetLastPing returns the last ping of the target
	GetLastPing() *Ping

	// NeedToPing returns true if the target needs to be pinged
	// based on the current time and last pinged time
	NeedToPing(cTime time.Time) bool
}

// GenericTarget is the generic type of target
type GenericTarget struct {
	// Unique identifier across database
	Id string `json:"id,omitempty"`

	// TargetType is the type of the target
	// Supported values are listed above
	TargetType string `json:"target_type,omitempty"`

	// Name of the target
	Name string `json:"name,omitempty"`

	// Creator of the target
	User *user.User `json:"-"`

	// Target details
	Protocol    string `json:"protocol"`
	HostAddress string `json:"host_address"`
	Port        int    `json:"port"`

	// PingInterval and PingTimeout is in seconds
	PingInterval_String string `json:"ping_interval"`
	PingInterval        int    `json:"-"`
	PingTimeout_String  string `json:"ping_timeout"`
	PingTimeout         int    `json:"-"`
	PingInProcess       bool   `json:"-"`
	LastPing            *Ping  `json:"-"`
}

// New returns a new Target instance
func New(gt *GenericTarget, usr *user.User) (Target, error) {
	// generic checks only
	// handled rest in each target type
	if gt == nil {
		return nil, errors.New("unable to get target details")
	}
	if gt.Name == "" {
		return nil, errors.New("name not specified")
	}
	gt.PingInterval = helpers.GetIntervalFromDurationStr(gt.PingInterval_String)
	gt.PingTimeout = helpers.GetIntervalFromDurationStr(gt.PingTimeout_String)
	if gt.PingInterval < 10 {
		return nil, errors.New("invalid ping_interval")
	}
	if gt.PingInterval > 86400 {
		return nil, errors.New("ping_interval cannot be more than a day")
	}
	if gt.PingTimeout < 1 {
		return nil, errors.New("invalid ping_timeout")
	}
	if gt.PingTimeout > gt.PingInterval {
		return nil, errors.New("ping_timeout cannot be more than ping_interval")
	}
	if gt.PingTimeout > 30 {
		return nil, errors.New("ping_timeout cannot be more than 30 seconds")
	}
	if usr == nil {
		return nil, errors.New("unable to get user details")
	}
	gt.User = usr
	gt.LastPing = nil
	gt.GenerateId()
	// to compare case insensitively
	gt.TargetType = strings.ToLower(gt.TargetType)

	switch gt.TargetType {
	case TargetType_Website:
		return NewWebsite(gt)
	default:
		return nil, fmt.Errorf("unsupported target type : %s", gt.TargetType)
	}
}

func (gt GenericTarget) GetGeneric() *GenericTarget {
	return &gt
}

func (gt GenericTarget) GetPoolKey() string {
	return Key(gt.Name, gt.User.Name)
}

func (gt *GenericTarget) GenerateId() {
	gt.Id = gt.GetPoolKey()
}

func (gt *GenericTarget) GetLastPing() *Ping {
	return gt.LastPing
}

func (gt *GenericTarget) NeedToPing(cTime time.Time) bool {
	if gt.LastPing == nil {
		return true
	}
	// PingInterval and Timestamp both are in seconds
	if (gt.LastPing.Timestamp+int64(gt.PingInterval)) < cTime.UnixMilli() && !gt.PingInProcess {
		return true
	}
	return false
}

func (gt *GenericTarget) pingInitiated() {
	gt.PingInProcess = true
}

func (gt *GenericTarget) pingDone() {
	gt.PingInProcess = false
}

// Key returns the key to use in the Pool
// currently it is of the format <userName>_<targetName>
func Key(userName, targetName string) string {
	return fmt.Sprintf(
		"%s_%s",
		userName, targetName,
	)
}

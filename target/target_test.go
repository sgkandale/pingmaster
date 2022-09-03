package target_test

import (
	"testing"

	"pingmaster/target"
	"pingmaster/user"
)

type GenericTargetTest struct {
	*target.GenericTarget
	ExpectError bool
}

func GetGenericTargestList() []GenericTargetTest {
	return []GenericTargetTest{
		// generic target checks
		{nil, true},
		{
			GenericTarget: &target.GenericTarget{},
			ExpectError:   true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name: "",
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 0,
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 86401,
			},
			ExpectError: true,
		},
		// website checks
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 20,
				TargetType:   target.TargetType_Website,
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 20,
				TargetType:   target.TargetType_Website,
				Protocol:     "tcp",
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 20,
				TargetType:   target.TargetType_Website,
				Protocol:     "https",
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 20,
				TargetType:   target.TargetType_Website,
				Protocol:     "https",
			},
			ExpectError: true,
		},
		{
			GenericTarget: &target.GenericTarget{
				Name:         "some",
				PingInterval: 20,
				TargetType:   target.TargetType_Website,
				Protocol:     "https",
				HostAddress:  "www.google.com",
			},
			ExpectError: false,
		},
	}
}

func TestTargetNew(t *testing.T) {
	targets := GetGenericTargestList()
	usr := &user.User{Name: "Ramesh"}
	for i := range targets {
		_, err := target.New(targets[i].GenericTarget, usr)
		if err != nil && !targets[i].ExpectError {
			t.Error(err)
		}
	}
}

package target

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Website struct {
	*GenericTarget
	Client *http.Client `json:"-"`
	URL    *url.URL     `json:"-"`
}

func NewWebsite(gt *GenericTarget) (Target, error) {
	if gt.TargetType != TargetType_Website {
		return nil, fmt.Errorf(
			"invalid target type in %s : %s",
			TargetType_Website, gt.TargetType,
		)
	}

	if gt.Protocol != "http" && gt.Protocol != "https" {
		return nil, fmt.Errorf(
			"unsupported protocol %s for target type %s",
			gt.Protocol, gt.TargetType,
		)
	}
	if gt.HostAddress == "" {
		return nil, errors.New("no host address provided for target type website")
	}
	u, err := url.Parse(gt.HostAddress)
	if err != nil {
		return nil, err
	}
	u.Scheme = gt.Protocol

	return &Website{
		GenericTarget: gt,
		Client:        http.DefaultClient,
		URL:           u,
	}, nil
}

func (w *Website) Ping(ctx context.Context) {
	w.pingInitiated()
	defer w.pingDone()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodHead,
		w.URL.String(),
		nil,
	)
	if err != nil {
		w.LastPing = &Ping{
			Timestamp: time.Now().Unix(),
			Error:     err,
		}
		return
	}

	startTime := time.Now().UnixMilli()
	resp, err := w.Client.Do(req)
	endTime := time.Now().UnixMilli()
	if err != nil {
		w.LastPing = &Ping{
			Timestamp: startTime / 1000,
			Duration:  int(endTime - startTime),
			Error:     err,
		}
		return
	}

	w.LastPing = &Ping{
		Timestamp:  startTime / 1000,
		Duration:   int(endTime - startTime),
		StatusCode: resp.StatusCode,
	}
}

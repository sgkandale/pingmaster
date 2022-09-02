package target_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"pingmaster/target"
)

func TestNewWebsite(t *testing.T) {
	_, err := target.NewWebsite(
		&target.GenericTarget{
			TargetType: target.TargetType_Website,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestWebsitePing(t *testing.T) {
	u, _ := url.Parse("https://www.google.com")
	ws := &target.Website{
		GenericTarget: &target.GenericTarget{
			TargetType: target.TargetType_Website,
		},
		Client: http.DefaultClient,
		URL:    u,
	}

	ws.Ping(context.Background())

	if ws.LastPing == nil {
		t.Error("last ping not generated")
		return
	}
	if ws.LastPing.Error != nil {
		t.Error(ws.LastPing.Error)
	}
}

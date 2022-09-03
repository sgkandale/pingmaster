package targetspool_test

import (
	"context"
	"testing"
	"time"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/target"
	"pingmaster/target/targetspool"
	"pingmaster/user"
)

var targetsList []target.Target

func init() {
	usr := &user.User{Name: "John"}
	tg, _ := target.New(
		&target.GenericTarget{
			TargetType:   target.TargetType_Website,
			Name:         "Google",
			Protocol:     "https",
			HostAddress:  "www.google.com",
			PingInterval: 10,
		},
		usr,
	)
	targetsList = []target.Target{tg}
}

func TestNewPool(t *testing.T) {
	pool := targetspool.New()
	if pool == nil {
		t.Error("NewPool returned nil")
		return
	}
	if pool.Targets == nil {
		t.Error("NewPool returned nil Targets")
	}
}

func TestPoolContains(t *testing.T) {
	pool := targetspool.New()

	gtArr := []target.Target{
		&target.Website{
			GenericTarget: &target.GenericTarget{
				Name: "Main",
				User: &user.User{
					Name: "Ramesh",
				},
			},
		},
	}

	for i := range gtArr {
		pool.Targets[gtArr[i].GetPoolKey()] = gtArr[i]
	}

	for i := range gtArr {
		if !pool.Contains(gtArr[i].GetPoolKey()) {
			t.Error("pool should contain given target")
		}
	}
}

func TestPoolGet(t *testing.T) {
	pool := targetspool.New()

	gtArr := []target.Target{
		&target.Website{
			GenericTarget: &target.GenericTarget{
				Name: "Main",
				User: &user.User{
					Name: "Ramesh",
				},
			},
		},
	}

	for i := range gtArr {
		pool.Targets[gtArr[i].GetPoolKey()] = gtArr[i]
	}

	for i := range gtArr {
		Tg := pool.Get(gtArr[i].GetPoolKey())
		if Tg == nil {
			t.Error("pool should contain given target")
			return
		}
		if Tg != gtArr[i] {
			t.Error("target not matching after adding")
		}
	}
}

func TestPoolAdd(t *testing.T) {
	pool := targetspool.New()

	gtArr := []target.Target{
		&target.Website{
			GenericTarget: &target.GenericTarget{
				Name: "Main",
				User: &user.User{
					Name: "Ramesh",
				},
			},
		},
	}

	for i := range gtArr {
		pool.Add(gtArr[i])
	}

	for i := range gtArr {
		Tg, tgExists := pool.Targets[gtArr[i].GetPoolKey()]
		if !tgExists || Tg == nil {
			t.Error("pool should contain given target")
		}
	}
}

func TestPoolMonitor(t *testing.T) {
	pool := targetspool.New()

	for i := range targetsList {
		pool.Add(targetsList[i])
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelCtx()
	dbConn, err := database.New(
		ctx,
		config.DatabaseConfig{},
	)
	if err != nil {
		t.Error(err)
		return
	}

	pool.Monitor(ctx, dbConn)
}

package target_test

import (
	"context"
	"testing"
	"time"

	"pingmaster/target"
	"pingmaster/user"
)

func TestNewPool(t *testing.T) {
	pool := target.NewPool()
	if pool == nil {
		t.Error("NewPool returned nil")
		return
	}
	if pool.Targets == nil {
		t.Error("NewPool returned nil Targets")
	}
}

func TestPoolContains(t *testing.T) {
	pool := target.NewPool()

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
	pool := target.NewPool()

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
	pool := target.NewPool()

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
	pool := target.NewPool()

	gtArr := GetTargetList()

	for i := range gtArr {
		pool.Add(gtArr[i])
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelCtx()

	pool.Monitor(ctx)
}

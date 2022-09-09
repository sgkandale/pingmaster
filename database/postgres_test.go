package database_test

import (
	"context"
	"log"
	"testing"
	"time"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/target"
	"pingmaster/user"

	"github.com/jackc/pgx/v4"
)

var (
	pgConfig = config.DatabaseConfig{
		Username:             "postgres",
		Password:             "postgres",
		Host:                 "localhost",
		Port:                 5432,
		DatabaseName:         "pingmaster",
		TimeoutInSeconds:     5,
		MaxConcurrentQueries: 100,
		PingsValidity:        30,
	}
)

func TestNewPostgres(t *testing.T) {

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}

	pgConn.Close(ctx)
}

func TestPostgresCheckUserExistance(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	name := "Ramesh"

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	userExist, errExists := pgConn.CheckUserExistance(
		ctx,
		user.User{
			Name: name,
		},
	)
	if errExists != nil && errExists != pgx.ErrNoRows {
		t.Error(errExists)
		return
	}

	if userExist {
		log.Println("[INF] user exists with name :", name)
	}
}

func TestPostgresGetUserDetails(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	name := "Ramesh"

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	usr := user.User{
		Name: name,
	}

	err = pgConn.GetUserDetails(ctx, &usr)
	if err != nil && err != pgx.ErrNoRows {
		t.Error(err)
	}
}

func TestPostgresInsertUser(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	name := "Ramesh"
	passwordHash := "abcd"

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	usr := user.User{
		Name:         name,
		PasswordHash: passwordHash,
	}

	err = pgConn.InsertUser(ctx, usr)
	if err != nil && err != pgx.ErrNoRows {
		t.Error(err)
	}
}

func TestPostgresFetchTargets(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	_, err = pgConn.FetchTargets(ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresInsertTarget(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	w, err := target.NewWebsite(
		&target.GenericTarget{
			TargetType: target.TargetType_Website,
			Name:       "Google",
			User: &user.User{
				Name: "Ramesh",
			},
			Protocol:     "https",
			HostAddress:  "www.google.com",
			PingInterval: 10,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = pgConn.InsertTarget(ctx, w)
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresInsertPing(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	ping := target.Ping{
		TargetKey:  "Google_Ramesh",
		Timestamp:  time.Now().Unix(),
		Duration:   150,
		StatusCode: 200,
		Error:      nil,
	}

	err = pgConn.InsertPing(ctx, ping)
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresDeleteOldPings(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	err = pgConn.DeleteOldPings(ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresGetTargetDetails(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	tg := target.GenericTarget{
		Name: "Google",
		User: &user.User{
			Name: "Ramesh",
		},
	}

	err = pgConn.GetTargetDetails(ctx, &tg)
	if err != nil && err != pgx.ErrNoRows {
		t.Error(err)
	}
}

func TestPostgresGetTargets(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	usr := user.User{
		Name: "Ramesh",
	}

	_, err = pgConn.GetTargets(ctx, usr)
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresGetPings(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	pgConn, err := database.NewPostgres(
		ctx,
		pgConfig,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer pgConn.Close(ctx)

	gt := target.GenericTarget{
		TargetType: target.TargetType_Website,
		Name:       "Google",
		User: &user.User{
			Name: "Ramesh",
		},
		Protocol:     "https",
		HostAddress:  "www.google.com",
		PingInterval: 10,
	}
	if err != nil {
		t.Error(err)
		return
	}

	_, err = pgConn.GetPings(ctx, &gt, time.Now().Unix(), 10)
	if err != nil {
		t.Error(err)
	}
}

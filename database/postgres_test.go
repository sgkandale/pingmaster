package database_test

import (
	"context"
	"log"
	"testing"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/target"
	"pingmaster/user"

	"github.com/jackc/pgx/v4"
)

var (
	pgConfig = config.DatabaseConfig{
		Username:         "postgres",
		Password:         "postgres",
		Host:             "localhost",
		Port:             5432,
		DatabaseName:     "pingmaster",
		TimeoutInSeconds: 5,
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

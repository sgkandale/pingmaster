package database_test

import (
	"context"
	"log"
	"testing"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/user"

	"github.com/jackc/pgx/v4"
)

var (
	pgConfig = config.DatabaseConfig{
		Username:         "postgres",
		Password:         "postgres",
		Host:             "localhost",
		Port:             5433,
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
		log.Println("[INFO] user exists with name :", name)
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
	name := "Rama"
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

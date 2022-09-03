package database

import (
	"context"
	"fmt"
	"strings"

	"pingmaster/config"
	"pingmaster/target"
	"pingmaster/user"
)

const (
	UsersTable   = "users"
	TargetsTable = "targets"
)

type Conn interface {
	Close(ctx context.Context)

	CheckUserExistance(ctx context.Context, usr user.User) (bool, error)

	GetUserDetails(ctx context.Context, usr *user.User) error

	InsertUser(ctx context.Context, usr user.User) error

	FetchTargets(ctx context.Context) ([]target.Target, error)

	InsertTarget(ctx context.Context, tg target.Target) error
}

func New(ctx context.Context, cfg config.DatabaseConfig) (Conn, error) {
	if strings.EqualFold(cfg.DatabaseType, "postgres") || strings.EqualFold(cfg.DatabaseType, "postgresql") {
		return NewPostgres(ctx, cfg)
	} else {
		return nil, fmt.Errorf("unsupported database type : %s", cfg.DatabaseType)
	}
}

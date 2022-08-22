package database

import (
	"context"
	"fmt"
	"strings"

	"pingmaster/config"
	"pingmaster/user"
)

const (
	UsersTable = "users"
)

type Conn interface {
	Close(ctx context.Context)

	CheckUserExistance(ctx context.Context, usr user.User) (bool, error)

	GetUserDetails(ctx context.Context, usr *user.User) error

	InsertUser(ctx context.Context, usr user.User) error
}

func New(ctx context.Context, cfg config.DatabaseConfig) (Conn, error) {
	if strings.EqualFold(cfg.DatabaseType, "postgres") || strings.EqualFold(cfg.DatabaseType, "postgresql") {
		return NewPostgres(ctx, cfg)
	} else {
		return nil, fmt.Errorf("unsupported database type : %s", cfg.DatabaseType)
	}
}

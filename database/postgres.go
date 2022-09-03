package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"pingmaster/config"
	"pingmaster/target"
	"pingmaster/user"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	DBConn  *pgxpool.Pool
	Timeout time.Duration
}

func NewPostgres(ctx context.Context, cfg config.DatabaseConfig) (Conn, error) {

	log.Printf(
		"[INF] connecting postgres at %s:%d",
		cfg.Host,
		cfg.Port,
	)

	dbPool, err := pgxpool.Connect(
		ctx,
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?connect_timeout=%d",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DatabaseName,
			cfg.TimeoutInSeconds,
		),
	)
	if err != nil {
		return nil, err
	}

	pingCtx, cancelPingCtx := context.WithTimeout(
		ctx,
		time.Second*time.Duration(cfg.TimeoutInSeconds),
	)
	defer cancelPingCtx()

	err = dbPool.Ping(pingCtx)
	if err != nil {
		return nil, err
	}

	log.Printf(
		"[INF] postgres connected at %s:%d",
		cfg.Host,
		cfg.Port,
	)

	return Postgres{
		DBConn:  dbPool,
		Timeout: time.Second * time.Duration(cfg.TimeoutInSeconds),
	}, nil
}

func (p Postgres) Close(ctx context.Context) {
	log.Println("[WRN] closing postgres conn")
	p.DBConn.Close()
}

func (p Postgres) CheckUserExistance(ctx context.Context, usr user.User) (bool, error) {

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()
	nameFromDb := ""

	err := p.DBConn.QueryRow(
		ctx,
		"select name from "+UsersTable+" where name = $1 LIMIT 1",
		usr.Name,
	).Scan(&nameFromDb)
	if err != nil && err != pgx.ErrNoRows {
		return true, err
	}
	if nameFromDb != "" {
		return true, nil
	}
	return false, nil
}

func (p Postgres) GetUserDetails(ctx context.Context, usr *user.User) error {

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	err := p.DBConn.QueryRow(
		ctx,
		"select password from "+UsersTable+" where name = $1 LIMIT 1",
		usr.Name,
	).Scan(&usr.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("not found")
		}
		return err
	}
	return nil
}

func (p Postgres) InsertUser(ctx context.Context, usr user.User) error {

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	_, err := p.DBConn.Exec(
		ctx,
		`insert into `+UsersTable+`
		 (name, password, created_at, last_login)
		 values ($1, $2, $3, $4);`,
		usr.Name,
		usr.PasswordHash,
		usr.CreatedAt,
		0,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) FetchTargets(ctx context.Context) ([]target.Target, error) {
	log.Printf("[INF] fetching targets from DB")

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	rows, err := p.DBConn.Query(
		ctx,
		`select key, name, type, creator, protocol, 
		host_address, port, ping_interval 
		from `+TargetsTable,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	tgs := []target.Target{}
	for rows.Next() {
		gt := target.GenericTarget{}
		usr := user.User{}
		err := rows.Scan(
			&gt.Id, &gt.Name, &gt.TargetType, &usr.Name, &gt.Protocol,
			&gt.HostAddress, &gt.Port, &gt.PingInterval,
		)
		if err != nil {
			log.Printf(
				"[ERR] scanning target from DB into struct : %s", err,
			)
			continue
		}
		tg, err := target.New(&gt, &usr)
		if err != nil {
			log.Printf(
				"[ERR] creating target out of data from DB : %s", err,
			)
			continue
		}
		tgs = append(tgs, tg)
	}
	log.Printf(
		"[INF] fetched %d targets from DB", len(tgs),
	)
	return tgs, nil
}

func (p Postgres) InsertTarget(ctx context.Context, tg target.Target) error {

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()
	gt := tg.GetGeneric()

	_, err := p.DBConn.Exec(
		ctx,
		`insert into `+TargetsTable+`
		 (key, name, type, creator, protocol,
			host_address, port, ping_interval)
		 values ($1, $2, $3, $4, $5, $6, $7, $8);`,
		gt.GetPoolKey(), gt.Name, gt.TargetType, gt.User.Name,
		gt.Protocol, gt.HostAddress, gt.Port, gt.PingInterval,
	)
	if err != nil {
		return err
	}
	return nil
}

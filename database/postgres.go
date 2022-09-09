package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"pingmaster/config"
	"pingmaster/helpers"
	"pingmaster/target"
	"pingmaster/user"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	DBConn               *pgxpool.Pool
	Timeout              time.Duration
	MaxConcurrentQueries int
	PingsValidity        int
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
		DBConn:               dbPool,
		Timeout:              time.Second * time.Duration(cfg.TimeoutInSeconds),
		MaxConcurrentQueries: cfg.MaxConcurrentQueries,
		PingsValidity:        cfg.PingsValidity,
	}, nil
}

func (p Postgres) GetMaxConcurrentQueries() int {
	return p.MaxConcurrentQueries
}
func (p Postgres) GetPingsValidity() int {
	return p.PingsValidity
}

func (p Postgres) Close(ctx context.Context) {
	log.Println("[WRN] closing postgres conn")
	p.DBConn.Close()
}

func (p Postgres) CheckUserExistance(ctx context.Context, usr user.User) (bool, error) {

	log.Printf(
		"[INF] checking user existance in DB : %s",
		usr.Name,
	)

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

	log.Printf(
		"[INF] getting user details from DB : %s",
		usr.Name,
	)

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

	log.Printf(
		"[INF] inserting user in DB : %s",
		usr.Name,
	)

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
		host_address, port, ping_interval, ping_timeout
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
			&gt.HostAddress, &gt.Port, &gt.PingInterval, &gt.PingTimeout,
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

	log.Printf(
		"[INF] inserting target in DB : %s %s:%d",
		gt.Protocol, gt.HostAddress, gt.Port,
	)

	_, err := p.DBConn.Exec(
		ctx,
		`insert into `+TargetsTable+`
		 (key, name, type, creator, protocol,
			host_address, port, ping_interval, ping_timeout)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
		gt.GetPoolKey(), gt.Name, gt.TargetType, gt.User.Name,
		gt.Protocol, gt.HostAddress, gt.Port, gt.PingInterval,
		gt.PingTimeout,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) InsertPing(ctx context.Context, ping target.Ping) error {

	log.Printf(
		"[INF] inserting ping in DB for : %s",
		ping.TargetKey,
	)

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	if ping.Error == nil {
		ping.Error = errors.New("")
	}

	_, err := p.DBConn.Exec(
		ctx,
		`insert into `+PingsTable+`
		 (key, timestamp, duration, status_code, error)
		 values ($1, $2, $3, $4, $5);`,
		ping.TargetKey, ping.Timestamp, ping.Duration,
		ping.StatusCode, ping.Error.Error(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) DeleteOldPings(ctx context.Context) error {

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()
	timestamp := time.Now().Add(time.Hour * -24 * time.Duration(p.PingsValidity)).Unix()

	log.Printf(
		"[INF] deleting old pings from DB before : %s",
		time.Unix(timestamp, 0).Format(helpers.Default_TimeFormat),
	)

	resp, err := p.DBConn.Exec(
		ctx,
		`delete from `+PingsTable+`
		where timestamp < $1;`,
		timestamp,
	)
	if err != nil {
		return err
	}
	log.Printf("[INF] old pings deleted : %+v", resp)

	return nil
}

func (p Postgres) GetTargetDetails(ctx context.Context, tg *target.GenericTarget) error {

	log.Printf(
		"[INF] getting target details from DB : %s",
		tg.GetPoolKey(),
	)

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	err := p.DBConn.QueryRow(
		ctx,
		`select key, name, type, protocol,
			host_address, port, ping_interval
			 from `+TargetsTable+` where key = $1 LIMIT 1`,
		tg.GetPoolKey(),
	).Scan(
		&tg.Id, &tg.Name, &tg.TargetType, &tg.Protocol,
		&tg.HostAddress, &tg.Port, &tg.PingInterval,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("not found")
		}
		return err
	}
	return nil
}

func (p Postgres) GetTargets(ctx context.Context, usr user.User) ([]target.Target, error) {

	log.Printf(
		"[INF] getting all targets from DB for user : %s",
		usr.Name,
	)

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	rows, err := p.DBConn.Query(
		ctx,
		`select key, name, type, protocol, 
		host_address, port, ping_interval, ping_timeout
		from `+TargetsTable+` where creator=$1`,
		usr.Name,
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
		err := rows.Scan(
			&gt.Id, &gt.Name, &gt.TargetType, &gt.Protocol,
			&gt.HostAddress, &gt.Port, &gt.PingInterval, &gt.PingTimeout,
		)
		if err != nil {
			log.Printf(
				"[ERR] scanning generic target from DB into struct : %s", err,
			)
			continue
		}
		ntg, err := target.New(&gt, &usr)
		if err != nil {
			log.Printf(
				"[ERR] creating target from generic target from DB : %s", err,
			)
			continue
		}
		tgs = append(tgs, ntg)
	}
	return tgs, nil
}

func (p Postgres) GetPings(ctx context.Context, gt *target.GenericTarget, beforeTime, limit int64) ([]*target.Ping, error) {

	log.Printf(
		"[INF] getting pings from DB for target : %s, before timestamp : %d, limit : %d",
		gt.GetPoolKey(), beforeTime, limit,
	)

	ctx, cancelCtx := context.WithTimeout(ctx, p.Timeout)
	defer cancelCtx()

	rows, err := p.DBConn.Query(
		ctx,
		`select timestamp, duration, status_code, error
		from `+PingsTable+` where key=$1 and
		timestamp<=$2::bigint order by timestamp desc limit $3::bigint;`,
		gt.GetPoolKey(), beforeTime, limit,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	pings := []*target.Ping{}
	for rows.Next() {
		ping := target.Ping{}
		pingErr := ""
		err := rows.Scan(
			&ping.Timestamp, &ping.Duration, &ping.StatusCode,
			&pingErr,
		)
		if err != nil {
			log.Printf(
				"[ERR] scanning ping from DB into struct : %s", err,
			)
			continue
		}
		ping.TimestampStr = time.Unix(ping.Timestamp, 0).Format(helpers.Default_TimeFormat)
		ping.Error = errors.New(pingErr)
		pings = append(pings, &ping)
	}
	return pings, nil
}

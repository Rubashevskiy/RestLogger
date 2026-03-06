package restlogger

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Heandler struct {
	repo *pgxpool.Pool
}

func NewHeandler(connString string) (*Heandler, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	} else {
		return &Heandler{repo: pool}, nil
	}
}


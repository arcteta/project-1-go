package db

import (
    "context"
    "fmt"
    "sync"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
    pool *pgxpool.Pool
}

var (
    pgInstance *Postgres
    pgOnce     sync.Once
)

func InitDBConnection(ctx context.Context, dbURL string) (*Postgres, error) {
    var err error

    pgOnce.Do(func() {
        pool, connErr := pgxpool.New(ctx, dbURL)
        if connErr != nil {
            err = fmt.Errorf("unable to create connection pool: %w", connErr)
            return
        }

        pgInstance = &Postgres{pool}
    })

    return pgInstance, err
}

func (pg *Postgres) Ping(ctx context.Context) error {
    return pg.pool.Ping(ctx)
}

func (pg *Postgres) Close() {
    pg.pool.Close()
}
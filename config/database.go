package config

import (
	"context"
	"fmt"
     db "github.com/furqanrupom/sqlc-crud/db"
	"github.com/jackc/pgx/v5/pgxpool"
)
func DBConnect(ctx context.Context,cfg DBConfig) (*pgxpool.Pool,*db.Queries,error) {
	
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",cfg.User,cfg.Password,cfg.Host,cfg.DBPort,cfg.DBName,cfg.SslEnableMode)
	pool,err := pgxpool.New(ctx,connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	pool.Config().MaxConns = 25
	pool.Config().MinConns = 5

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("pool ping failed: %w", err)
	}

	queries := db.New(pool) 

	return pool, queries, nil


}


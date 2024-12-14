package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/microservice-name/config"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgresConnection(cfg *config.DatabaseConfig) (*Postgres, error) {
	// Construct connection string
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	// Configure connection pool
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %v", err)
	}

	// Set pool configuration
	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MinConns = 1
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	// Create connection pool with context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func (p *Postgres) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.Pool.Ping(ctx)
}

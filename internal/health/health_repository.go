package health

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthRepository struct {
	db *pgxpool.Pool
}

func NewHealthRepository(db *pgxpool.Pool) *HealthRepository {
	return &HealthRepository{db: db}
}

func (r *HealthRepository) CheckDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.db.Ping(ctx)
}

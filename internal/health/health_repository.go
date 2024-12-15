package health

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/department-service-api/internal/errors"
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
	if err := r.db.Ping(ctx); err != nil {
		return errors.NewDatabaseConnectionError(err)
	}
	return nil
}

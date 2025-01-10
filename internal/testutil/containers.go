// internal/testutil/container.go
package testutil

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestContainer struct {
	Container testcontainers.Container
	Pool      *pgxpool.Pool
}

func SetupTestContainer(ctx context.Context) (*TestContainer, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:14-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections"),
			wait.ForListeningPort("5432/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %v", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %v", err)
	}

	connString := fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable&pool_max_conns=10&pool_max_conn_lifetime=1h",
		hostIP, mappedPort.Port())

	var db *pgxpool.Pool
	for i := 0; i < 5; i++ {
		dbConfig, err := pgxpool.ParseConfig(connString)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		db, err = pgxpool.ConnectConfig(ctx, dbConfig)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		if err := db.Ping(ctx); err != nil {
			db.Close()
			time.Sleep(time.Second)
			continue
		}

		break
	}

	if db == nil {
		return nil, fmt.Errorf("failed to connect to database after retries")
	}

	return &TestContainer{
		Container: container,
		Pool:      db,
	}, nil
}

func (tc *TestContainer) Teardown(ctx context.Context) error {
	if tc.Pool != nil {
		tc.Pool.Close()
	}

	if tc.Container != nil {
		return tc.Container.Terminate(ctx)
	}
	return nil
}

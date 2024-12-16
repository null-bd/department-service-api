package health

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestContainer(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()

	// Define PostgreSQL container configuration
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForAll(
			wait.NewLogStrategy("database system is ready to accept connections").
				WithStartupTimeout(30*time.Second),
			wait.ForListeningPort("5432/tcp"),
		),
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		AutoRemove: true,
	}

	// Start container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	// Get container connection details
	mappedPort, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	hostIP, err := container.Host(ctx)
	require.NoError(t, err)

	// Build connection string
	connString := "postgres://test:test@" + hostIP + ":" + mappedPort.Port() + "/testdb?sslmode=disable"

	// Wait a bit to ensure PostgreSQL is fully ready
	time.Sleep(2 * time.Second)

	return container, connString
}

func SetupDatabase(t *testing.T, connString string) *pgxpool.Pool {
	ctx := context.Background()

	// Configure connection pool
	config, err := pgxpool.ParseConfig(connString)
	require.NoError(t, err)

	// Set reasonable pool settings
	config.MaxConns = 5
	config.MinConns = 1
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute

	// Connect with retry logic
	var pool *pgxpool.Pool
	for i := 0; i < 3; i++ {
		pool, err = pgxpool.ConnectConfig(ctx, config)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	require.NoError(t, err)

	// Verify connection
	err = pool.Ping(ctx)
	require.NoError(t, err)

	return pool
}

package internal_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/dan-cook/valkey-investigation/internal"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCacheWithTestContainers(t *testing.T) {
	ctx := context.Background()

	con, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		// repeat what we would have in docker compose
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "valkey/valkey:7.2.5",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("Ready to accept connections tcp"), // fragile and really needs a timeout but works...
		},
		Started: true,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		con.Terminate(ctx)
	})
	cache, err := internal.NewValKey()
	require.NoError(t, err)

	t.Cleanup(cache.Close)

	err = cache.Ping()
	require.NoError(t, err)

	key := fmt.Sprintf("key-%d", time.Now().UnixNano())

	err = cache.Set(key, "value")
	require.NoError(t, err)

	value, err := cache.Get(key)
	require.NoError(t, err)

	require.Equal(t, "value", value)
}

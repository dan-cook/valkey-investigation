package internal_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dan-cook/valkey-investigation/internal"
	"github.com/stretchr/testify/require"
)

// TestCache tests the cache implementation can set and get a value
func TestCache(t *testing.T) {
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

package ratelimiter

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uhonliu/go-utils"
)

func TestRedisRatelimiter(t *testing.T) {
	rdb, err := goutils.NewRedisClient(&redis.Options{})
	require.Nil(t, err)
	limiter := NewRedisRatelimiter(rdb)
	ctx := context.Background()
	assert.True(t, limiter.Allow(ctx, "key", time.Second*1, 1))
	assert.False(t, limiter.Allow(ctx, "key", time.Second*1, 1))
	time.Sleep(1 * time.Second)
	assert.True(t, limiter.Allow(ctx, "key", time.Second*1, 1))
}

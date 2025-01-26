package kv

import (
	"context"
	"time"
)

type KV interface {
	Add(ctx context.Context, k string, v string, ttl time.Duration) error
	Get(ctx context.Context, k string) (string, error)
	Exist(ctx context.Context, k string) (bool, error)
}

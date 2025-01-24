package kv

import "context"

type KV interface {
	Add(ctx context.Context, k string) error
	Exist(ctx context.Context, k string) (bool, error)
}

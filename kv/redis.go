package kv

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisKVConfig struct {
	URL       string        `env:"REDIS_URL"`
	KeyPrefix string        `env:"REDIS_KEY_PREFIX"`
	ttl       time.Duration `env:"REDIS_TTL"`
}
type RedisKV struct {
	client *redis.Client
	prefix string
	ttl    time.Duration
}

func NewRedisKV(c RedisKVConfig) *RedisKV {
	opt, err := redis.ParseURL(c.URL)
	if err != nil {
		log.Fatalf("Error parsing redis url: %v", err)
	}
	client := redis.NewClient(opt)
	return &RedisKV{
		client: client,
		prefix: c.KeyPrefix,
		ttl:    c.ttl,
	}
}

func (r RedisKV) getKey(k string) string {
	return fmt.Sprintf("%s:%s", r.prefix, k)
}

func (r RedisKV) Add(ctx context.Context, k string) error {
	return r.client.Set(ctx, r.getKey(k), 1, r.ttl).Err()
}

func (r RedisKV) Exist(ctx context.Context, k string) (bool, error) {
	return r.client.Exists(ctx, r.getKey(k)).Val() == 1, nil
}

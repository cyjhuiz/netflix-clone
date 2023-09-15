package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDao struct {
	client *redis.Client
}

func NewRedisDao() *RedisDao {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
	return &RedisDao{client: redisClient}
}

func (redisDao *RedisDao) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := redisDao.client.Get(ctx, key).Result()
	if err != nil {
		return value, err
	}
	return value, nil
}

func (redisDao *RedisDao) Set(key string, value string) error {
	ctx := context.Background()
	defaultCacheExpiry := 12 * time.Hour
	_, err := redisDao.client.Set(ctx, key, value, defaultCacheExpiry).Result()
	if err != nil {
		return err
	}
	return nil
}

func (redisDao *RedisDao) Delete(key string) error {
	ctx := context.Background()
	_, err := redisDao.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (redisDao *RedisDao) Flush() error {
	ctx := context.Background()
	_, err := redisDao.client.FlushDB(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

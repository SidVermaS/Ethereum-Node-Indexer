package configs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Client *redis.Client

func CreateCacheConnection(wg *sync.WaitGroup) {
	defer wg.Done()
	var addr = fmt.Sprintf("%s:%s", os.Getenv(string(consts.REDIS_HOST)), os.Getenv(string(consts.REDIS_PORT)))
	var options *redis.Options = &redis.Options{
		Addr:     addr,
		Password: os.Getenv(string(consts.REDIS_PASSWORD)),
		DB:       0,
	}
	Client = redis.NewClient(options)
}

func SetCacheValue(key consts.CacheE, value interface{}) error {
	err := Client.Set(ctx, string(key), value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
func SetCacheValueWithExpiry(key consts.CacheE, value interface{}, expiry time.Duration) error {
	err := Client.Set(ctx, string(key), value, 0).Err()
	if err != nil {
		return err
	}
	_, err1 := Client.Expire(ctx, string(key), expiry).Result()
	if err1 != nil {
		return err1
	}
	return nil
}
func GetCacheValue(key consts.CacheE) (string, error) {
	value, err := Client.Get(ctx, string(key)).Result()

	if err != nil {
		return "", err
	}
	return value, nil
}
func DeleteCacheValue(key consts.CacheE) (int64, error) {
	value, err := Client.Del(ctx, string(key)).Result()
	if err != nil {
		return -1, err
	}
	return value, nil
}

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

// Connecting to the Cache
func CreateCacheConnection(wg *sync.WaitGroup) {
	defer wg.Done()
	// Address to the redis cache server
	var addr = fmt.Sprintf("%s:%s", os.Getenv(string(consts.REDIS_HOST)), os.Getenv(string(consts.REDIS_PORT)))

	//
	var options *redis.Options = &redis.Options{
		// host:port address.
		Addr:     addr,
		Password: os.Getenv(string(consts.REDIS_PASSWORD)),
		// Database to be selected after connecting to the server.
		DB: 0,
	}
	// It returns a client to the Redis Server specified by Options
	Client = redis.NewClient(options)
}

func CloseCacheConnection(wg *sync.WaitGroup) {
	defer wg.Done()
	Client.Close()
}

// It sets the key and values in the cache
func SetCacheValue(key consts.CacheE, value interface{}) error {
	err := Client.Set(ctx, string(key), value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// It sets the key and values in the cache with an expiry time
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

// It fetches the value by the key passed which was saved with either the SetCacheValue() or SetCacheValueWithExpiry()
func GetCacheValue(key consts.CacheE) (string, error) {
	value, err := Client.Get(ctx, string(key)).Result()

	if err != nil {
		return "", err
	}
	return value, nil
}

// Delete a value by key which was saved
func DeleteCacheValue(key consts.CacheE) (int64, error) {
	value, err := Client.Del(ctx, string(key)).Result()
	if err != nil {
		return -1, err
	}
	return value, nil
}

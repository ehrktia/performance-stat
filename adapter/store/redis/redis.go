package redis

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v9"
)

const (
	redis_timeout = 15 * time.Second
	defaultHost   = "0.0.0.0"
	defaultPort   = "6379"
	HOST          = "HOST"
	PORT          = "PORT"
)

type redisStore struct {
	rcStore *redis.Client
}

func isNotValid(in string) bool {
	return strings.EqualFold(in, "")

}

func getURLFromEnv() string {
	var h, p string
	if h = os.Getenv(HOST); isNotValid(h) {
		h = defaultHost
	}
	if p = os.Getenv(PORT); isNotValid(p) {
		p = defaultPort
	}
	return fmt.Sprintf("%s:%s", h, p)
}

func New(ctx context.Context) (*redisStore, error) {
	url:=getURLFromEnv()
	redisClient := redis.NewClient(&redis.Options{
		Addr:        url,
		Password:    "",
		DialTimeout: redis_timeout,
		ReadTimeout: redis_timeout,
	})
	out, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", out)
	return &redisStore{
		rcStore: redisClient,
	}, nil
}

package redis

import (
	"os"

	"github.com/go-redis/redis"
)

var Client = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_URI"),
})

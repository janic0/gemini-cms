package redis

import "github.com/go-redis/redis"

var Client = redis.NewClient(&redis.Options{})

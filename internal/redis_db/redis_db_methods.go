package redis_db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redis_db_connection *redis.Client = Redis_db_connection()

func Get(key_name string) string {
	val, err := redis_db_connection.Get(ctx, key_name).Result()
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func Set(key_name string, value string) bool {
	err := redis_db_connection.Set(ctx, key_name, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

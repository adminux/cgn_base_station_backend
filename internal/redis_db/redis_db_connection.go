package redis_db

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func Redis_db_connection() *redis.Client {
	redis_json_config, _ := os.ReadFile("../../configs/redis_connection_settings.json")
	var redis_db_conn_settings RedisDbConnSettings
	err := json.Unmarshal(redis_json_config, &redis_db_conn_settings)
	if err != nil {
		log.Fatal(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     redis_db_conn_settings.Host + ":" + strconv.Itoa(redis_db_conn_settings.Port),
		Password: "",
		DB:       redis_db_conn_settings.Db,
	})
}

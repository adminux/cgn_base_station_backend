package main

import (
	"fmt"
	"os"

	"github.com/adminux/cgn_base_station_backend/internal/redis_db"
)

func main() {
	path, exists := os.LookupEnv("GOPATH")
	if exists {
		fmt.Println("___GOPATH = ", path)
	}

	path, exists = os.LookupEnv("GOBIN")
	if exists {
		fmt.Println("___GOBIN = ", path)
	}

	path, exists = os.LookupEnv("PATH")
	if exists {
		fmt.Println("___PATH = ", path)
	}

	redis_db.Redis_db_connection()
}

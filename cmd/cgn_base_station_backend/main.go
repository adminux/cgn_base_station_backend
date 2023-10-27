package main

import (
	"fmt"

	"github.com/adminux/cgn_base_station_backend/internal/redis_db"
)

func main() { 
	redis_db.Redis_db_connection()
	fmt.Println("GPS:hostname = ", redis_db.Get("GPS:hostname"))
}

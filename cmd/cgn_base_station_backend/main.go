package main

import (
	"github.com/adminux/cgn_base_station_backend/internal/redis_db"
	"github.com/adminux/cgn_base_station_backend/internal/server"
)

func main() {
	redis_db.Redis_db_connection()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}

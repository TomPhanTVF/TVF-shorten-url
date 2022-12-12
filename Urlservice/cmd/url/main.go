package main

import (
	"url-service/internal/server"
	"url-service/internal/sql"
)

func main() {
	sql := &sql.Sql{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "postgres",
		DbName:   "url_db",
	}
	sql.ConnectDB()
	defer sql.Close()
	server.RunServerRPC()
}


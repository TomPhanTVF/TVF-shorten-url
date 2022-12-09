package main

import (
	"user-service/internal/sql"
	"user-service/internal/server"

)

func main(){
	sql := &sql.Sql{
		Host:     "localhost",
		Port:      "5432",
		Username: "postgres",
		Password: "postgres",
		DbName:   "auth_db",
	}
	sql.ConnectDB()
	defer sql.Close()
	server.RunServerRPC()
}
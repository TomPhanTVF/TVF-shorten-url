package main

import (
	"log"
	"user-service/internal/server"
	"user-service/internal/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	sql := &sql.Sql{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DbName:   "auth_db",
	}
	sql.ConnectDB()
	defer sql.Close()
	// run migration
	// RunMiration("postgresql://postgres:postgres@localhost:5432/auth_db?sslmode=disable", "file://internal/migration")
	// run gRPC server
	server.RunServerRPC()
}

func RunMiration(dbMigrationURL string, dbSource string) {
	migration, err := migrate.New(dbMigrationURL, dbSource)
	if err != nil {
		log.Fatal("can not create new migrate intance", err)
	}
	if err = migration.Up(); err != nil {
		log.Fatal("failed to run migrate up:", err)
	}
	log.Println("Migrate successfuly")
}

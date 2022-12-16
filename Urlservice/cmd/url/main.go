package main

import (
	"url-service/internal/server"
	pb"url-service/pb/url"
	"url-service/internal/sql"
	"url-service/internal/handle"
	pg "url-service/internal/repository/postgres"
	cli "url-service/client/user"
)

func main() {
	sql := &sql.Sql{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "postgres",
		DbName:   "shorten_db",
	}
	sql.ConnectDB()
	defer sql.Close()
	urlRepo := pg.NewURlRepo(sql.Db)
	userClient := cli.NewUserClientService()
	urlServer := handle.NewURLServer(pb.UnimplementedURLServiceServer{}, userClient, urlRepo)
	serverPRC := server.NewServerRPC(urlServer)
	serverPRC.RunServerRPC()
	
}


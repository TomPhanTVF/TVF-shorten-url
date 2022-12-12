package user

import (
	pb "url-service/pb/user"
	"google.golang.org/grpc"
	"log"
)



type UserClient struct {
	UserService  *pb.UserServiceClient
}

func InitProductServiceClient(url string) *UserClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		log.Println("Could not connect:", err)
	}

	c := &UserClient{
		UserService: pb.NewUserServiceClient(),

	}

	return c
}
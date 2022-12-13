package user

import (
	"context"
	pb "url-service/pb/user"
	"google.golang.org/grpc"
	"log"
)


type UserClientService struct {
}


func NewUserClientService()UserClientService{
	return UserClientService{}
}



func(u *UserClientService)FindUserByEmail(email string) *pb.FindByEmailResponse{
	conn, err := grpc.DialContext(context.Background(), ":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	
	req := &pb.FindByEmailRequest{
		Email: email,
	}
	res, err := client.FindByEmail(context.Background(), req)
	if err != nil {
		log.Printf("call FindByEmail err %v\n", err)
		return nil
	}
	return &pb.FindByEmailResponse{
		User: &pb.User{
			Id: res.User.Id,
			FirstName: res.User.FirstName,
			LastName: res.User.LastName,
			Email: res.User.Email,
			Role: res.User.Role,
		},
	}
}
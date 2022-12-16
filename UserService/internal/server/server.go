package server

import(
	"net"
	"google.golang.org/grpc"
	"user-service/pb"
	"user-service/internal/handle"
	"log"
	interceptor "user-service/internal/user_interceptor"
)


// Run service
func RunServerRPC() error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}
	interceptor := interceptor.NewAuthInterceptor(accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}
	s := grpc.NewServer(serverOptions...)
	pb.RegisterUserServiceServer(s, &handle.AuthServiceHandle{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/TVF-shorten-url.pcuser.UserService/"

	return map[string][]string{
		laptopServicePath + "FindByID": {"admin","user"},
		laptopServicePath + "FindByEmail":  {"admin","user"},
		laptopServicePath + "GetMe":   {"admin", "user"},
	}
}

package server

import(
	"net"
	"google.golang.org/grpc"
	"user-service/pb"
	"user-service/internal/handle"
	"log"
)


// Run service
func RunServerRPC() error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handle.AuthServiceHandle{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}
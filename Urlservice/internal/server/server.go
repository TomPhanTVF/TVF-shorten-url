package server

import(
	"net"
	"google.golang.org/grpc"
	pb"url-service/pb/url"
	"url-service/internal/handle"
	"log"
)


type severRPC struct{
	handle handle.URLServer
}

func NewServerRPC(handle handle.URLServer) *severRPC{
	return &severRPC{
		handle: handle,
	}
}

// Run service
func(server severRPC) RunServerRPC() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Unable to listen on port 4000: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterURLServiceServer(s, &server.handle)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return nil
}
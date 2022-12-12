package server

import(
	"net"
	"google.golang.org/grpc"
	pb"url-service/pb/url"
	"url-service/internal/handle"
	"log"
)


// Run service
func RunServerRPC() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterURLServiceServer(s, &handle.URLServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return nil
}
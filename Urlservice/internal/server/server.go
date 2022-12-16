package server

import(
	"net"
	"google.golang.org/grpc"
	pb"url-service/pb/url"
	"url-service/internal/handle"
	interceptor "url-service/internal/url_interceptor"
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
	interceptor := interceptor.NewAuthInterceptor(accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}
	s := grpc.NewServer(serverOptions...)
	pb.RegisterURLServiceServer(s, &server.handle)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return nil
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/TVF-shorten-url.pcurl.URLService/"

	return map[string][]string{
		laptopServicePath + "CreateUrl": {"user"},
		laptopServicePath + "GetURLsByOwner":  {"user"},
		laptopServicePath + "GetURLs":   {"admin", "user"},
		laptopServicePath + "DeleteURl":   { "user"},
	}
}
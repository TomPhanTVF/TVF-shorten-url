package handle

import (
	pb "url-service/pb/url"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"context"
)



type URLServer struct {
	pb.UnimplementedURLServiceServer
}


func (u *URLServer) CreateUrl(context.Context, *pb.CreateUrlReq) (*pb.CreateUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (u *URLServer) DeleteURl(context.Context, *pb.DeleteUrlReq) (*pb.DeleteUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteURl not implemented")
}
func (u *URLServer) GetURLsByOwnerID(context.Context, *pb.GetURLsByOwnerIDReq) (*pb.GetURLsByOwnerIDReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLsByOwnerID not implemented")
}
func (u *URLServer) GetURLs(context.Context, *pb.GetURLsReq) (*pb.GetURLsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLs not implemented")
}
func (u *URLServer) GetRedirectByTVF(context.Context, *pb.GetRedirectByTVFReq) (*pb.GetRedirectByTVFRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedirectByTVF not implemented")
}
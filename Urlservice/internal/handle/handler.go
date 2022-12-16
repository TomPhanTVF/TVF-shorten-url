package handle

import (
	pb "url-service/pb/url"
	Upb "url-service/pb/user"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"context"
	"url-service/client/user"
	"url-service/internal/utils"
	"url-service/internal/security"
	"url-service/internal/models"
	url "url-service/internal/repository/postgres"
	_ "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)



type URLServer struct {
	pb.UnimplementedURLServiceServer
	user.UserClientService
	url.URLRepo
}

func NewURLServer(unimplement pb.UnimplementedURLServiceServer, userClient user.UserClientService, urlRepo url.URLRepo) URLServer{
	return URLServer{
		unimplement,
		userClient,
		urlRepo,
	}
}


func (u *URLServer) CreateUrl(ctx context.Context, in *pb.CreateUrlReq) (*pb.CreateUrlRes, error) {
	token, err := u.getTokenromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "ParseToken: %v", err)
	}
	claim, err := security.Verify(token)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "VerifyToken: %v", err)
	}
	user := u.FindUserByEmail(claim.Email)
	createUrl := &models.URL{
		Redirect : in.GetRedirect(),
		TVF      :in.GetTVF(),
		UserID   :user.User.GetId(),
	}
	createUrl.GenID()
	createUrl.PrepareBeforeInsert()


	url, err := u.Create(ctx, createUrl)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "CreateURl: %v", err)
	}
	createURLRes := &pb.CreateUrlRes{
		Id       :url.Id,
		Redirect :url.Redirect,
		TVF      :url.TVF,
		Random   :url.Random,
	}

	return createURLRes, nil
}
func (u *URLServer) DeleteURl(context.Context, *pb.DeleteUrlReq) (*pb.DeleteUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteURl not implemented")
}
func (u *URLServer) GetURLsByOwner(ctx context.Context, in *pb.GetURLsByOwnerReq) (*pb.GetURLsByOwnerRes, error) {
	token, err := u.getTokenromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "ParseToken: %v", err)
	}
	claim, err := security.Verify(token)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "VerifyToken: %v", err)
	}
	user := u.FindUserByEmail(claim.Email)
	urls, err := u.FindUrlsByEmail(ctx, user.User.Id)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "GetURLsByOwner: %v", err)
	}
	var urlsRPC  []*pb.URL
	for _, v := range urls{
		urlMapping := &pb.URL{
			Id: v.Id,
			Redirect: v.Redirect,
			TVF: v.TVF,
			Random: v.Random,
			Owner: &Upb.User{
				Id        :user.User.Id,
				FirstName :user.User.FirstName,
				LastName  :user.User.LastName,
				Password  :user.User.Password,
				Email     :user.User.Email,
				Role      :user.User.Role,
			},
		}
		urlsRPC = append(urlsRPC, urlMapping)
	}

	return &pb.GetURLsByOwnerRes{
		Urls: urlsRPC,
	}, nil
}	
func (u *URLServer) GetURLs(context.Context, *pb.GetURLsReq) (*pb.GetURLsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLs not implemented")
}
func (u *URLServer) GetRedirectByTVF(ctx context.Context, in *pb.GetRedirectByTVFReq) (*pb.GetRedirectByTVFRes, error) {
	redirect, err := u.URLRepo.GetRedirectByTVF(ctx, in.TVF)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "GetRedirectByTVF: %v", err)
	}
	redirectRPCResponse := &pb.GetRedirectByTVFRes{
		Redirect: redirect,
	}

	return redirectRPCResponse, nil
}

func (u *URLServer) getTokenromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext: %v", utils.ErrNoCtxMetaData)
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]

	return accessToken, nil
}


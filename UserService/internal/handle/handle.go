package handle

import (
	"context"
	models "user-service/internal/models"
	"user-service/internal/security"
	"user-service/internal/service"
	"user-service/internal/utils"
	"user-service/pb"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type AuthServiceHandle struct {
	pb.UnimplementedUserServiceServer
	userSV	service.UserService
}



func (a *AuthServiceHandle) Register(ctx context.Context, input *pb.RegisterRequest) (*pb.RegisterResponse, error){
	user := &models.User{
	FirstName : input.GetFirstName(),
	LastName  : input.GetLastName(),
	Password  : input.GetPassword(),
	Email     : input.GetEmail(),
	Role      : input.GetRole(),
	}
	if err := utils.ValidateStruct(ctx, user); err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "ValidateStruct: %v", err)
	}
	createdUser, err := a.userSV.Register(ctx, user)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "Register: %v", err)
	}
	userResponsePRC := &pb.RegisterResponse{
		User: &pb.User{
			Id        :createdUser.ID,
			FirstName :createdUser.FirstName,
			LastName  :createdUser.LastName,
			Password  :createdUser.Password,
			Email     :createdUser.Email,
			Role      :createdUser.Role,

		},
	}

	return userResponsePRC, nil
}
func (a *AuthServiceHandle) FindByEmail(ctx context.Context, input *pb.FindByEmailRequest) (*pb.FindByEmailResponse, error){
	email := input.GetEmail()
	if !utils.ValidateEmail(email) {
		return nil, status.Errorf(codes.InvalidArgument, "ValidateEmail: %v", email)
	}

	user, err := a.userSV.FindByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "userUC.FindByEmail: %v", err)
	}
	userResponsePRC := &pb.FindByEmailResponse{
		User: &pb.User{
			Id        :user.ID,
			FirstName :user.FirstName,
			LastName  :user.LastName,
			Password  :user.Password,
			Email     :user.Email,
			Role      :user.Role,

		},
	}
	return userResponsePRC, nil
}
func (a *AuthServiceHandle) FindByID(ctx context.Context, input *pb.FindByIDRequest) (*pb.FindByIDResponse, error){
	userUUID:= input.GetUuid()
	
	user, err := a.userSV.FindById(ctx, userUUID)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "userUC.FindById: %v", err)
	}
	userResponsePRC := &pb.FindByIDResponse{
		User: &pb.User{
			Id        :user.ID,
			FirstName :user.FirstName,
			LastName  :user.LastName,
			Password  :user.Password,
			Email     :user.Email,
			Role      :user.Role,
		},
	}
	return userResponsePRC, nil
}
func (a *AuthServiceHandle) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error){
	email := input.GetEmail()
	if !utils.ValidateEmail(email) {

		return nil, status.Errorf(codes.InvalidArgument, "ValidateEmail: %v", email)
	}

	user, err := a.userSV.Login(ctx, email, input.GetPassword())
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "Login: %v", err)
	}
	token, err  := security.Gentoken(user)
	loginRes := &pb.LoginResponse{
		User: &pb.User{
			Id        :user.ID,
			FirstName :user.FirstName,
			LastName  :user.LastName,
			Password  :user.Password,
			Email     :user.Email,
			Role      :user.Role,
		},
		Token: token,	
	}

	return loginRes, nil
}
func (a *AuthServiceHandle) GetMe(ctx context.Context, input *pb.GetMeRequest) (*pb.GetMeResponse, error){
	token, err := a.getTokenromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	claim, err := security.Verify(token)
	if err != nil {
		return nil, err
	}
	user, err := a.userSV.FindByEmail(ctx, claim.Email)
	if err != nil {
		return nil, status.Errorf(utils.ParseGRPCErrStatusCode(err), "userUC.GetMe: %v", err)
	}
	userResponsePRC := &pb.GetMeResponse{
		User: &pb.User{
			Id        :user.ID,
			FirstName :user.FirstName,
			LastName  :user.LastName,
			Password  :user.Password,
			Email     :user.Email,
			Role      :user.Role,
		},
	}
	return userResponsePRC, nil
}
func (a *AuthServiceHandle) Logout(ctx context.Context, input *pb.LogoutRequest) (*pb.LogoutResponse, error){
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext: %v", utils.ErrNoCtxMetaData)
	}
	delete(md, "authorization")

	return &pb.LogoutResponse{}, nil
}

func (a *AuthServiceHandle) getTokenromCtx(ctx context.Context) (string, error) {
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

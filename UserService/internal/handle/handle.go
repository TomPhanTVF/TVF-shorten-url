package handle

import (
	"context"
	models "user-service/internal/models"
	"user-service/internal/service"
	"user-service/internal/utils"
	"user-service/pb"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
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
func (a *AuthServiceHandle) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error){
	return nil, nil
}
func (a *AuthServiceHandle) GetMe(context.Context, *pb.GetMeRequest) (*pb.GetMeResponse, error){
	return nil, nil
}
func (a *AuthServiceHandle) Logout(context.Context, *pb.LogoutRequest) (*pb.LogoutResponse, error){
	return nil, nil
}

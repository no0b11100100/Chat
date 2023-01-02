package services

import (
	"Chat/Server/api"
	log "Chat/Server/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	api.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SignIn(_ context.Context, userData *api.SignIn) (*api.Response, error) {
	log.Info.Printf("SignIn %+v\n", *userData)
	var result api.Response
	return &result, nil
}

func (s *UserService) SignUp(_ context.Context, userData *api.SignUp) (*api.Response, error) {
	log.Info.Printf("SignUp %+v\n", *userData)

	if userData.Password != userData.ConfirmedPassword {
		return &api.Response{StatusMessage: "Passwords not match"}, nil
	}
	var result api.Response
	return &result, nil
}

func (s *UserService) EditUser(context.Context, *api.UserData) (*empty.Empty, error) {
	return nil, nil
}

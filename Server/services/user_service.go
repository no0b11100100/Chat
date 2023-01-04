package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	api.UnimplementedUserServer
	database interfaces.UserServiceDatabase
}

func NewUserService(database interfaces.UserServiceDatabase) *UserService {
	return &UserService{database: database}
}

func (s *UserService) SignIn(_ context.Context, userData *api.SignIn) (response *api.Response, err error) {
	log.Info.Printf("SignIn %+v\n", *userData)

	status, userID := s.database.ValidateUser(userData.Email, userData.Password)
	if !status {
		response.Status = api.ResponseStatus_OK
		return
	}

	response.UserId = userID

	return
}

func (s *UserService) SignUp(_ context.Context, userData *api.SignUp) (response *api.Response, err error) {
	log.Info.Printf("SignUp %+v\n", *userData)

	if userData.Password != userData.ConfirmedPassword {
		response.Status = api.ResponseStatus_OK
		response.StatusMessage = "Passwords not match"
		return
	}

	if !s.database.IsEmailUnique(userData.Email) {
		response.Status = api.ResponseStatus_OK
		response.StatusMessage = "Email already in use"
		return
	}

	_, userID := s.database.RegisterUser(userData)
	s.database.AddUserToChat(userID, "-1") //Just for test: "-1" is a test chat

	response.UserId = userID

	return
}

func (s *UserService) EditUser(context.Context, *api.UserData) (*empty.Empty, error) {
	return nil, nil
}

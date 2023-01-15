package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
)

type UserService struct {
	database interfaces.UserServiceDatabase
}

func NewUserService(database interfaces.UserServiceDatabase) *UserService {
	return &UserService{database: database}
}

func (s *UserService) SignIn(_ api.ServerContext, userData api.SignIn) api.Response {
	log.Info.Printf("SignIn %+v\n", userData)
	response := api.Response{Info: api.UserInfo{}}
	status, userID := s.database.ValidateUser(userData.Email, userData.Password)
	if !status {
		response.Status = api.OK
		return response
	}

	response.Info.UserID = userID

	return response
}
func (s *UserService) SignUp(_ api.ServerContext, userData api.SignUp) api.Response {
	log.Info.Printf("SignUp %+v\n", userData)
	response := api.Response{Info: api.UserInfo{}}
	if userData.Password != userData.ConfirmedPassword {
		response.Status = api.OK
		response.StatusMessage = "Passwords not match"
		return response
	}

	if !s.database.IsEmailUnique(userData.Email) {
		response.Status = api.OK
		response.StatusMessage = "Email already in use"
		return response
	}

	_, userID := s.database.RegisterUser(userData)
	s.database.AddUserToChat(userID, "-1") //Just for test: "-1" is a test chat

	response.Info.UserID = userID

	return response
}

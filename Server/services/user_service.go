package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
)

type UserService struct {
	database interfaces.UserServiceDatabase
	users    map[string]string
}

func NewUserService(serviceConnection *ServiceConnection, database interfaces.UserServiceDatabase) *UserService {
	s := &UserService{database: database, users: make(map[string]string)}
	serviceConnection.ConnectionIDByUserEmail.Provide(s.connectionIDByUserID)
	return s
}

func (s *UserService) SignIn(ctx api.ServerContext, userData api.SignIn) api.Response {
	log.Info.Printf("SignIn %+v\n", userData)
	response := api.Response{Info: api.UserInfo{Chats: make([]string, 0)}}
	status, userID := s.database.ValidateUser(userData.Email, userData.Password)
	if !status {
		response.Status = api.OK
		return response
	}

	response.Info.UserID = userID
	s.saveUser(userData.Email, ctx.ConnectionID)

	return response
}

func (s *UserService) SignUp(ctx api.ServerContext, userData api.SignUp) api.Response {
	log.Info.Printf("SignUp %+v\n", userData)
	response := api.Response{Info: api.UserInfo{Chats: make([]string, 0)}}
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
	s.saveUser(userData.Email, ctx.ConnectionID)

	return response
}

func (s *UserService) saveUser(email, connectionID string) {
	s.users[email] = connectionID
}

func (s *UserService) connectionIDByUserID(userID string) string {
	return s.users[userID]
}

func (s *UserService) HandleDisconnect(connectionID string) {
	for key, value := range s.users {
		if value == connectionID {
			delete(s.users, key)
			return
		}
	}
}

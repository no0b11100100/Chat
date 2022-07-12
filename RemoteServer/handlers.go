package main

import (
	"Chat/RemoteServer/common"
	"encoding/json"
	"fmt"
	"log"
)

func (s *Server) SignIn(payload []byte) common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.SignInOK, Type: common.SignIn}}

	if err != nil {
		log.Println(err)
		response.Command.Status = common.ServerError
		return response
	}

	status, userID := s.database.ValidateUser(user)
	if !status {
		response.Command.Status = common.SignInError
		return response
	}

	response.Command.Payload = []byte(string(fmt.Sprintf("{\"user_id\":\"%v\"}", userID)))
	return response
}

func (s *Server) SignUp(payload []byte) common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.SignUpOK, Type: common.SignUp}}

	if err != nil {
		log.Println(err)
		response.Command.Status = common.ServerError
		return response
	}

	if !s.database.IsEmailUnique(user.Email) {
		response.Command.Status = common.SignUpInvalidEmail
		return response
	}

	_, userID := s.database.RegisterUser(user)

	response.Command.Payload = []byte(string(fmt.Sprintf("{\"user_id\":\"%v\"}", userID)))
	return response
}

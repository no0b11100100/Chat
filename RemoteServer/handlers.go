package main

import (
	"Chat/RemoteServer/database"
	"common"
	"encoding/json"
	"fmt"
	"log"
)

func (s *Server) LogIn(payload []byte) common.CommandResponce {
	info := common.UserInfo{}
	err := json.Unmarshal(payload, &info)
	if err != nil {
		log.Println(err)
		return common.CommandResponce{Error: "Invalid input"}
	}

	user := database.User{Email: info.Email, Password: info.Password}
	status, userID := s.database.ValidateUser(user)
	if !status {
		return common.CommandResponce{Error: "LogIn failed: Invalid data"}
	}

	return common.CommandResponce{Command: common.Command{Type: common.LogIn, Payload: []byte(string(fmt.Sprintf("{\"id\":%v}", userID)))}}
}

func (s *Server) RegisterUser(payload []byte) common.CommandResponce {
	info := common.UserInfo{}
	err := json.Unmarshal(payload, &info)
	if err != nil {
		log.Println(err)
		return common.CommandResponce{Error: "Invalid input"}
	}

	user := database.User{Email: info.Email, Password: info.Password}
	userID := s.database.RegisterUser(user)

	return common.CommandResponce{Command: common.Command{Type: common.Register, Payload: []byte(string(fmt.Sprintf("{\"id\":%v}", userID)))}}
}

func (s *Server) SendMessage([]byte) common.CommandResponce { return common.CommandResponce{} }

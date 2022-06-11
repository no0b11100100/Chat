package main

import (
	"common"
	"encoding/json"
	"fmt"
	"log"
)

func (s *Server) LogIn(payload []byte) common.CommandResponce {
	// user := common.User{}
	// err := json.Unmarshal(payload, &user)
	// if err != nil {
	// 	log.Println(err)
	// 	return common.CommandResponce{Error: "Invalid input"}
	// }

	// status, userID := s.database.ValidateUser(user)
	// if !status {
	// 	return common.CommandResponce{Error: "LogIn failed: Invalid data"}
	// }

	// return common.CommandResponce{Command: common.Command{Type: common.LogIn, Payload: []byte(string(fmt.Sprintf("{\"id\":%v}", userID)))}}

	return common.CommandResponce{Command: common.Command{Type: common.LogIn, Payload: []byte(string(fmt.Sprintf("{\"id\":\"%v\"}", "ID from remote server")))}}
}

func (s *Server) RegisterUser(payload []byte) common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	if err != nil {
		log.Println(err)
		return common.CommandResponce{Error: "Invalid input"}
	}

	if s.database.IsEmailUnique(user.Email) {
		ok, userID := s.database.RegisterUser(user)
		if !ok {
			return common.CommandResponce{Error: "Server error"}
		}
		return common.CommandResponce{Command: common.Command{Type: common.Register, Payload: []byte(string(fmt.Sprintf("{\"id\":%v}", userID)))}}
	}

	return common.CommandResponce{Error: "Email is already used"}
}

func (s *Server) SendMessage(payload []byte) common.CommandResponce {
	message := common.Message{}
	err := json.Unmarshal(payload, &message)
	if err != nil {
		log.Println(err)
		return common.CommandResponce{Error: "Invalid input"}
	}

	s.database.AddMessage(message)

	return common.CommandResponce{}
}

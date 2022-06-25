package main

import (
	"common"
	"encoding/json"
	"fmt"
	"log"
)

func (s *Server) SignIn(payload []byte) common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	response := common.CommandResponce{Type: common.Response, common.Command{Status: common.OK, Type: common.SignIn}}

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

	response.Payload = []byte(string(fmt.Sprintf("{\"id\":\"%v\"}", userID)))
	return response
}

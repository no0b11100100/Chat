package main

import (
	"Chat/RemoteServer/common"
	log "Chat/RemoteServer/common/logger"
	api "Chat/RemoteServer/structs"
	"encoding/json"
	"fmt"
)

func (s *Server) SignIn(payload []byte) *common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.SignInOK, Type: common.SignIn}}

	if err != nil {
		log.Warning.Println(err)
		response.Command.Status = common.ServerError
		return &response
	}

	status, userID := s.database.ValidateUser(user)
	if !status {
		response.Command.Status = common.SignInError
		return &response
	}

	response.Command.Payload = []byte(string(fmt.Sprintf("{\"user_id\":\"%v\"}", userID)))
	return &response
}

func (s *Server) SignUp(payload []byte) *common.CommandResponce {
	user := common.User{}
	err := json.Unmarshal(payload, &user)
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.SignUpOK, Type: common.SignUp}}

	if err != nil {
		log.Warning.Println(err)
		response.Command.Status = common.ServerError
		return &response
	}

	if !s.database.IsEmailUnique(user.Email) {
		response.Command.Status = common.SignUpInvalidEmail
		return &response
	}

	_, userID := s.database.RegisterUser(user)

	response.Command.Payload = []byte(string(fmt.Sprintf("{\"user_id\":\"%v\"}", userID)))
	return &response
}

func (s *Server) GetUserChats(payload []byte) *common.CommandResponce {
	log.Info.Println("GetUserChats", payload)
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.OK, Type: common.GetUserChatsCommand}}
	userID := api.UserID{}

	err := json.Unmarshal(payload, &userID)
	if err != nil {
		log.Warning.Println(err)
	}

	chats := s.database.GetUserChats(userID.UserId)

	p, err := json.Marshal(chats)
	if err != nil {
		log.Warning.Println(err)
	}

	response.Command.Payload = p

	return &response
}

func (s *Server) GetMessages(payload []byte) *common.CommandResponce {
	response := common.CommandResponce{Type: common.Response, Command: common.Command{Status: common.OK, Type: common.GetMessagesCommand}}
	messageChan := api.MessageChan{}
	err := json.Unmarshal(payload, &messageChan)
	if err != nil {
		log.Warning.Println(err)
	}

	messages := s.database.GetMessages(messageChan.ChatId, messageChan.MessageId, messageChan.Direction)
	p, err := json.Marshal(messages)
	if err != nil {
		log.Warning.Println(err)
	}
	response.Command.Payload = p

	return &response
}

func (s *Server) SendMessage(payload []byte) *common.CommandResponce {
	log.Info.Println("SendMessage", payload)
	//TODO: get chatID
	s.broadcastChat("chatID", payload)
	return nil
}

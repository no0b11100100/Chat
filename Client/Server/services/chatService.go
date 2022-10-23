package services

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	log "Chat/Client/Server/common/logger"
	"Chat/Client/Server/interfaces"
	"context"
	"encoding/json"

	"github.com/golang/protobuf/ptypes/empty"
)

type chatService struct {
	sender interfaces.Sender
	api.UnimplementedChatServer
	messageNotification chan api.ExchangedMessage
}

func NewChatService(sender interfaces.Sender) *chatService {
	return &chatService{sender: sender, messageNotification: make(chan api.ExchangedMessage, 1)}
}

func (chat *chatService) IsHandledTopic(command common.CommandType) bool {
	return command == common.NotifyMessageCommand
}

func (chat *chatService) HandleTopic(c common.Command) {
	switch c.Type {
	case common.NotifyMessageCommand:
		log.Info.Println("Get Message", string(c.Payload))
		message := api.ExchangedMessage{}
		if err := json.Unmarshal(c.Payload, &message); err != nil {
			log.Error.Println(err)
		}
		chat.messageNotification <- message
	}
}

func (chat *chatService) GetChats(_ context.Context, userID *api.UserID) (*api.Chats, error) {
	var chats api.Chats

	c := common.Command{Type: common.GetUserChatsCommand}
	var err error
	c.Payload, err = json.Marshal(*userID)

	if err != nil {
		log.Warning.Println(err)
	}
	response := <-chat.sender.Send(c)
	log.Info.Println(response)
	err = json.Unmarshal(response.Payload, &chats)
	if err != nil {
		log.Warning.Println(err)
	}
	return &chats, nil
}

func (chat *chatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	var messages []*api.Message
	c := common.Command{Type: common.GetMessagesCommand}
	p, _ := json.Marshal(*messageChan)
	c.Payload = p

	response := <-chat.sender.Send(c)

	err := json.Unmarshal(response.Payload, &messages)
	if err != nil {
		log.Warning.Println(err)
	}

	result := &api.Messages{Messages: messages}
	return result, nil
}

func (chat *chatService) ReadMessage(context.Context, *api.ReadMessage) (*api.Status, error) {
	return nil, nil
}

func (chat *chatService) EditChat(context.Context, *api.ChatData) (*api.Status, error) {
	return nil, nil
}

func (chat *chatService) ChatChanged(*empty.Empty, api.Chat_ChatChangedServer) error {
	return nil
}

func (chat *chatService) SendMessage(_ context.Context, message *api.ExchangedMessage) (*api.Status, error) {
	log.Info.Printf("SendMessage %+v\n", *message)
	c := common.Command{Type: common.SendMessageCommand}
	var err error
	c.Payload, err = json.Marshal(*message)

	if err != nil {
		log.Warning.Println(err)
	}
	for data := range chat.sender.Send(c) {
		log.Info.Println("SendMessage response", data)
	}
	return &api.Status{Status: "OK"}, nil
}

func (chat *chatService) RecieveMessage(_ *empty.Empty, stream api.Chat_RecieveMessageServer) error {
	log.Info.Println("RecieveMessage")

	for message := range chat.messageNotification {
		log.Info.Println("ReciveMessage", message)
		if err := stream.Send(&message); err != nil {
			log.Error.Println(err)
		}
	}

	return nil
}

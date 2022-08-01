package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	log "Chat/Client/Server/common/logger"
	"context"
	"encoding/json"
)

type ChatService struct {
	sender RemoteServerInterface
	api.UnimplementedChatServer
}

func NewChatService(sender RemoteServerInterface) *ChatService {
	return &ChatService{sender: sender}
}

type ResponseType interface {
	api.UserChats | api.ChatInformation | api.ParticipantInfo | api.Messages
}

func makeRequest[T ResponseType](send func(common.Command, common.ChannelType), command common.CommandType, payload interface{}, response *T) {
	c := common.Command{Type: common.GetUserChatsCommand}
	var err error
	c.Payload, err = json.Marshal(payload)

	if err != nil {
		log.Warning.Println(err)
	}

	ch := make(common.ChannelType)
	defer close(ch)
	send(c, ch)

	result := <-ch

	err = json.Unmarshal(result.Payload, &response)
	if err != nil {
		log.Warning.Println(err)
	}
}

func (chat *ChatService) GetUserChats(_ context.Context, userID *api.UserID) (*api.UserChats, error) {
	log.Info.Printf("GetUserChats %+v\n", *userID)
	var chats api.UserChats
	makeRequest(chat.sender.Send, common.GetUserChatsCommand, *userID, &chats)
	return &chats, nil
}

func (chat *ChatService) GetChatInfo(_ context.Context, chatID *api.ChatID) (*api.ChatInformation, error) {
	log.Info.Printf("GetChatInfo %+v\n", *chatID)
	var chatInfo api.ChatInformation
	makeRequest(chat.sender.Send, common.GetChatInfoCommand, *chatID, &chatInfo)
	return &chatInfo, nil
}

func (chat *ChatService) GetParticipantInfo(_ context.Context, userID *api.UserID) (*api.ParticipantInfo, error) {
	log.Info.Printf("GetParticipantInfo %+v\n", *userID)
	var participantInfo api.ParticipantInfo
	makeRequest(chat.sender.Send, common.GetParticipantInfoCommand, *userID, &participantInfo)
	return &participantInfo, nil
}

func (chat *ChatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	var messages api.Messages
	makeRequest(chat.sender.Send, common.GetMessagesCommand, *messageChan, &messages)
	return &messages, nil
}

func (chat *ChatService) MessageExchange(api.Chat_MessageExchangeServer) error {
	return nil
}

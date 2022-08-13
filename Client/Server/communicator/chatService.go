package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	log "Chat/Client/Server/common/logger"
	"context"
	"encoding/json"

	"github.com/golang/protobuf/ptypes/empty"
)

type ChatService struct {
	sender RemoteServerInterface
	api.UnimplementedChatServer
}

func NewChatService(sender RemoteServerInterface) *ChatService {
	return &ChatService{sender: sender}
}

type ResponseType interface {
	api.Messages | api.Chats
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

func (chat *ChatService) GetChats(_ context.Context, userID *api.UserID) (*api.Chats, error) {
	log.Info.Printf("GetChats %+v\n", *userID)
	var chats api.Chats
	var c api.ChatInfo
	c.ChatId = "1"
	c.Title = "ALice"
	c.SecondLine = ""
	c.LastMessage = "test message"
	c.UnreadedCount = 0
	c.Cover = ""
	chats.Chats = append(chats.Chats, &c)
	// makeRequest(chat.sender.Send, common.GetUserChatsCommand, *userID, &chats)
	return &chats, nil
}

func (chat *ChatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	var messages api.Messages
	messages.Messages = append(messages.Messages, &api.Message{MessageJson: string([]byte(`{"text":"test message"}`))})
	// makeRequest(chat.sender.Send, common.GetMessagesCommand, *messageChan, &messages)
	return &messages, nil
}

func (chat *ChatService) ReadMessage(context.Context, *api.ReadMessage) (*empty.Empty, error) {
	return nil, nil
}

func (chat *ChatService) EditChat(context.Context, *api.ChatData) (*empty.Empty, error) {
	return nil, nil
}

func (chat *ChatService) ChatChanged(*empty.Empty, api.Chat_ChatChangedServer) error {
	return nil
}

func (chat *ChatService) SendMessage(context.Context, *api.ExchangedMessage) (*empty.Empty, error) {
	return nil, nil
}

func (chat *ChatService) RecieveMessage(*empty.Empty, api.Chat_RecieveMessageServer) error {
	return nil
}

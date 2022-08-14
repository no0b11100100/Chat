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
	messageChan chan *api.ExchangedMessage
}

func NewChatService(sender RemoteServerInterface) *ChatService {
	return &ChatService{sender: sender, messageChan: make(chan *api.ExchangedMessage)}
}

type ResponseType interface {
	[]*api.Message | api.Chats | empty.Empty
}

func makeRequest[T ResponseType](send func(common.Command, common.ChannelType), command common.CommandType, payload interface{}, response *T) {
	c := common.Command{Type: command}
	var err error
	c.Payload, err = json.Marshal(payload)

	if err != nil {
		log.Warning.Println(err)
	}

	if response == nil {
		send(c, nil)
	} else {
		ch := make(common.ChannelType)
		defer close(ch)
		send(c, ch)

		result := <-ch

		err = json.Unmarshal(result.Payload, &response)
		if err != nil {
			log.Warning.Println(err)
		}
	}
}

func (chat *ChatService) GetChats(_ context.Context, userID *api.UserID) (*api.Chats, error) {
	log.Info.Printf("GetChats %+v\n", *userID)
	var chats api.Chats
	// var c api.ChatInfo
	// c.ChatId = "1"
	// c.Title = "ALice"
	// c.SecondLine = ""
	// c.LastMessage = "test message"
	// c.UnreadedCount = 0
	// c.Cover = ""
	// chats.Chats = append(chats.Chats, &c)
	makeRequest(chat.sender.Send, common.GetUserChatsCommand, *userID, &chats)
	return &chats, nil
}

func (chat *ChatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	var messages []*api.Message
	// messages.Messages = append(messages.Messages, &api.Message{MessageJson: string([]byte(`{"text":"test message"}`))})
	makeRequest(chat.sender.Send, common.GetMessagesCommand, *messageChan, &messages)
	result := &api.Messages{Messages: messages}
	return result, nil
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

func (chat *ChatService) SendMessage(_ context.Context, message *api.ExchangedMessage) (*empty.Empty, error) {
	log.Info.Printf("SendMessage %+v\n", *message)
	// chat.messageChan <- message
	c := common.Command{Type: common.SendMessageCommand}
	var err error
	c.Payload, err = json.Marshal(*message)

	if err != nil {
		log.Warning.Println(err)
	}
	chat.sender.Send(c, nil)

	return &empty.Empty{}, nil
}

func (chat *ChatService) RecieveMessage(_ *empty.Empty, stream api.Chat_RecieveMessageServer) error {
	for {
		for message := range chat.messageChan {
			err := stream.Send(message)
			if err != nil {
				log.Warning.Println(err)
			}
		}
	}

	return nil
}

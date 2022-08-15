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

func NewChatService(sender RemoteServerInterface, notification common.ChannelType) *ChatService {
	chatService := &ChatService{sender: sender, messageChan: make(chan *api.ExchangedMessage)}

	go chatService.handleNotification(notification)

	return chatService
}

func (chat *ChatService) handleNotification(ch common.ChannelType) {
	for {
		for data := range ch {
			if data.Type == common.NotifyMessageCommand {
				var message api.ExchangedMessage
				err := json.Unmarshal(data.Payload, &message)
				if err == nil {
					chat.messageChan <- &message
				} else {
					log.Warning.Println(string(data.Payload), err)
				}
			}
		}
	}
}

func (chat *ChatService) IsHandledTopic(topic common.CommandType) bool {
	return topic == common.NotifyMessageCommand
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

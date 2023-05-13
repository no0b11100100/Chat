package database

import (
	api "Chat/Server/api"
	log "Chat/Server/logger"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatsDatabase struct {
	*Base
}

func NewChatsDatabase() *ChatsDatabase {
	return &ChatsDatabase{&Base{}}
}

func (db *ChatsDatabase) Connect(client *mongo.Client, database string) {
	db.Base.Connect(client, database, "Chats")
}

func (db *ChatsDatabase) GetChatInfo(chatID string) api.Chat {
	log.Info.Println(chatID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var chat api.Chat
	chatTagger := api.ChatTags()
	err := db.Collection.FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	return chat
}

func (db *ChatsDatabase) AddMessage(msg api.Message) error {
	chatID := msg.ChatID

	fmt.Println("AddMessage", chatID)
	var chat api.Chat
	chatTagger := api.ChatTags()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return err
	}

	text := struct {
		Text string `json:"text"`
	}{}

	json.Unmarshal(msg.MessageJSON, &text)

	chat.Messages = append(chat.Messages, msg)
	chat.LastMessage = text.Text

	filter := bson.D{{chatTagger.ChatID, chatID}}
	update := bson.D{{"$set", bson.D{{chatTagger.Messages, chat.Messages}, {chatTagger.LastMessage, chat.LastMessage}}}}
	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return err
	}

	return nil
}

func (db *ChatsDatabase) AddUserToChat(userID string, chatID string) {
	var chat api.Chat
	chatTagger := api.ChatTags()
	fmt.Println("AddUserToChat tag", chatTagger.ChatID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	chat.Participants = append(chat.Participants, userID)
	filter := bson.D{{chatTagger.ChatID, chatID}}
	update := bson.D{{"$set", bson.D{{chatTagger.Participants, chat.Participants}}}}
	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

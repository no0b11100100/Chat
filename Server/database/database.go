package database

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"

	api "Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
)

// //https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

type DB struct {
	m      sync.RWMutex
	client *mongo.Client
	tables map[string]*mongo.Collection
}

func NewDatabase() *DB {
	return &DB{
		m:      sync.RWMutex{},
		tables: make(map[string]*mongo.Collection),
	}
}

type Database interface {
	interfaces.UserServiceDatabase
	interfaces.ChatServiceDatabase
	Connect()
	Close()
}

func (db *DB) Connect() {
	fmt.Println("Connect database")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/db")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Info.Println("DB Connect error:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Info.Println("DB Connect ping error:", err)
	}

	log.Info.Println("Connected to MongoDB!")

	db.client = client

	users := client.Database("application").Collection("Users")
	chats := client.Database("application").Collection("Chats")

	db.tables["Users"] = users
	db.tables["Chats"] = chats

	// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	// _, _ = db.tables["Chats"].InsertOne(ctx, api.Chat{
	// 	Messages: []api.Message{
	// 		{
	// 			MessageJSON: []byte(`{"text":"test message"}`),
	// 		},
	// 	},
	// 	ChatID:       "-1",
	// 	Title:        "Test chat",
	// 	LastMessage:  "test message",
	// 	Participants: []string{},
	// })
	// if err != nil {
	// 	log.Warning.Println(err)
	// }

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Error.Fatal(err)
	}
	log.Info.Println(databases)
}

func (db *DB) Close() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.client.Disconnect(ctx)

	if err != nil {
		log.Error.Println("DB Close error:", err)
	}
	log.Info.Println("Connection to MongoDB closed.")
}

func (db *DB) IsEmailUnique(email string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	userTagger := api.UserInfoTags()
	err := db.tables["Users"].FindOne(ctx, bson.M{userTagger.Email: email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true
		}
	}
	return false
}

func (db *DB) ValidateUser(email, password string) (bool, string) {
	log.Info.Println("Validate user", email, password)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var record bson.M
	userTagger := api.UserInfoTags()
	if err := db.tables["Users"].FindOne(ctx, bson.M{userTagger.Email: email, userTagger.Password: password}).Decode(&record); err != nil {
		log.Warning.Println(err)
		return false, ""
	}
	log.Info.Println(record)

	userID, ok := record["userid"].(string)
	if !ok {
		log.Warning.Println("ValidateUser error")
		return false, ""
	}

	return true, userID
}

func (db *DB) RegisterUser(data api.SignUp) (bool, string) {
	user := api.UserInfo{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
		NickName: data.NickName,
		Photo:    data.Photo,
	}

	user.UserID = uuid.New().String()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.tables["Users"].InsertOne(ctx, user)
	if err != nil {
		log.Warning.Println(err)
		return false, ""
	}

	return true, user.UserID
}

func (db *DB) AddUserToChat(userID string, chatID string) {
	var chat api.Chat
	chatTagger := api.ChatTags()
	fmt.Println("AddUserToChat tag", chatTagger.ChatID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.tables["Chats"].FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	chat.Participants = append(chat.Participants, userID)
	filter := bson.D{{chatTagger.ChatID, chatID}}
	update := bson.D{{"$set", bson.D{{chatTagger.Participants, chat.Participants}}}}
	_, err = db.tables["Chats"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	var user api.UserInfo
	userTagger := api.UserInfoTags()
	if err := db.tables["Users"].FindOne(ctx, bson.M{userTagger.UserID: userID}).Decode(&user); err != nil {
		log.Warning.Println(err)
		return
	}

	user.Chats = append(user.Chats, chatID)
	filter = bson.D{{userTagger.UserID, userID}}
	update = bson.D{{"$set", bson.D{{userTagger.Chats, user.Chats}}}}
	_, err = db.tables["Users"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

func (db *DB) getChatInfo(chatID string) api.Chat {
	log.Info.Println(chatID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var chat api.Chat
	chatTagger := api.ChatTags()
	err := db.tables["Chats"].FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	return chat
}

func (db *DB) GetUserChats(userID string) []api.Chat {
	log.Info.Println("GetUserChatsr", userID)
	result := make([]api.Chat, 0)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var user api.UserInfo
	userTagger := api.UserInfoTags()

	err := db.tables["Users"].FindOne(ctx, bson.M{userTagger.UserID: userID}).Decode(&user)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	for _, chatID := range user.Chats {
		chat := db.getChatInfo(chatID)
		result = append(result, chat)
	}

	log.Info.Println(result)

	return result
}

// // If message_from is empty - return all available messages for chat
// // otherwise - messages from provided message_id
func (db *DB) GetMessages(chatID string) []api.Message {
	log.Info.Printf("GetMessages %v\n", chatID)
	chat := db.getChatInfo(chatID)

	return chat.Messages
}

func (db *DB) AddMessage(msg api.Message) error {
	chatID := msg.ChatID

	fmt.Println("AddMessage", chatID)
	var chat api.Chat
	chatTagger := api.ChatTags()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.tables["Chats"].FindOne(ctx, bson.M{chatTagger.ChatID: chatID}).Decode(&chat)
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
	_, err = db.tables["Chats"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return err
	}

	return nil
}

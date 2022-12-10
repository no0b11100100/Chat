package database

import (
	"Chat/RemoteServer/common"
	log "Chat/RemoteServer/common/logger"
	api "Chat/RemoteServer/structs"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

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
	Connect()
	Close()
	IsEmailUnique(string) bool
	ValidateUser(common.User) (bool, string)
	RegisterUser(common.User) (bool, string)
	GetUserChats(string) api.Chats
	// GetChatParticipants(string) []api.ParticipantInfo
	GetMessages(string, string, api.Direction) []*api.Message
	// GetUsers() []common.User
	// GetChatMessages(string) []common.Message
	// AddMessage(common.Message)
	// AddChat(common.Chat) error
	// AddUserToChat(string) error
	// LeaveChat(string, string) error
	AddUserToChat(userID string, chatID string)
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
	// _, _ = db.tables["Chats"].InsertOne(ctx, api.ChatInfo{Messages: &api.Messages{Messages: []*api.Message{&api.Message{MessageJson: string([]byte(`{"message":"test message"}`))}}},
	// 	ChatId: "1", Title: "Test chat", LastMessage: "test message", Participants: []string{"726ac197-8640-43cc-9f54-0006780957f1"}})
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
	err := db.tables["Users"].FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true
		}
	}
	return false
}

func (db *DB) ValidateUser(user common.User) (bool, string) {
	log.Info.Println("Validate user", user)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var record bson.M
	if err := db.tables["Users"].FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&record); err != nil {
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

func (db *DB) RegisterUser(userInfo common.User) (bool, string) {
	var user api.UserInfo

	user.Name = userInfo.Name
	user.NickName = userInfo.NickName
	user.Email = userInfo.Email
	user.Password = userInfo.Password

	user.UserId = uuid.New().String()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.tables["Users"].InsertOne(ctx, user)
	if err != nil {
		log.Warning.Println(err)
		return false, ""
	}

	return true, user.UserId
}

// func (db *DB) CreateChat(isStorage bool, creator string, members []string) {
// 	var chat api.ChatInfo // TODO: fill data

// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	db.tables["Chats"].InsertOne(ctx, chat)
// }

func (db *DB) AddUserToChat(userID string, chatID string) {
	var chat api.ChatInfo
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.tables["Chats"].FindOne(ctx, bson.M{"chatId": chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	chat.Participants = append(chat.Participants, userID)
	filter := bson.D{{"chatId", chatID}}
	update := bson.D{{"$set", bson.D{{"participants", chat.Participants}}}}
	_, err = db.tables["Chats"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	var user api.UserInfo
	if err := db.tables["Users"].FindOne(ctx, bson.M{"userid": userID}).Decode(&user); err != nil {
		log.Warning.Println(err)
		return
	}

	user.Chats = append(user.Chats, chatID)
	filter = bson.D{{"userid", userID}}
	update = bson.D{{"$set", bson.D{{"chats", user.Chats}}}}
	_, err = db.tables["Users"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

func (db *DB) getChatInfo(chatID string) api.ChatInfo {
	log.Info.Println(chatID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var chat api.ChatInfo
	err := db.tables["Chats"].FindOne(ctx, bson.M{"chatId": chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	return chat
}

func (db *DB) GetUserChats(userID string) api.Chats {
	log.Info.Println("GetUserChatsr", userID)
	result := make([]*api.ChatInfo, 0)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var user api.UserInfo

	err := db.tables["Users"].FindOne(ctx, bson.M{"userid": userID}).Decode(&user)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	for _, chatID := range user.Chats {
		chat := db.getChatInfo(chatID)
		result = append(result, &chat)
	}

	log.Info.Println(result)

	return api.Chats{Chats: result}
}

// If message_from is empty - return all available messages for chat
// otherwise - messages from provided message_id
func (db *DB) GetMessages(chatID string, message_from string, direction api.Direction) []*api.Message {
	log.Info.Printf("GetMessages %v < %v >\n", chatID, message_from)
	chat := db.getChatInfo(chatID)

	return chat.Messages.Messages
	// result := make([]api.Message, 0)
	// if chatID == "1" {
	// 	result = append(result, api.Message{MessageJson: string([]byte(`{"sender":"Alice","message":"test message", "sender_id":"726ac197-8640-43cc-9f54-0006780957f1"}`))})
	// }
	// if message_from == "" {
	// 	var chat api.ChatInfo
	// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 	err := db.tables["Chats"].FindOne(ctx, bson.M{"chat_id": chatID}).Decode(&chat)
	// 	if err != nil {
	// 		log.Warning.Println("DB error:", err)
	// 	}

	// 	// return chat.Messages
	// } else {
	// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 	err := db.tables["Chats"].FindOne(ctx, bson.M{"chat_id": chatID}) // $elementMatch : {message_id : { $gt : message_from } }
	// 	if err != nil {
	// 		log.Warning.Println("DB error:", err)
	// 	}
	// }

	// return result
}

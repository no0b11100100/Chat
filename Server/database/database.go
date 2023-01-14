package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "Chat/Server/logger"
)

// import (
// 	api "Chat/Server/api"
// 	interfaces "Chat/Server/interfaces"
// 	log "Chat/Server/logger"
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/google/uuid"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

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
	// interfaces.UserServiceDatabase
	// interfaces.ChatServiceDatabase
	Connect()
	Close()
	AddUserToChat(userID string, chatID string)
}

// type UserInfo struct {
// 	UserID   string
// 	Email    string
// 	Password string
// 	Name     string
// 	NickName string
// 	Photo    string
// }

// type ChatInfo struct {
// 	ChatID        string
// 	Title         string
// 	SecondLine    string
// 	LastMessage   string
// 	UnreadedCount int
// 	Cover         string
// 	Participants  []string
// 	Messages      []api.Message
// }

// func (db *DB) createTestChat() {
// 	var chat ChatInfo
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	err := db.tables["Chats"].FindOne(ctx, bson.M{"chatid": "-1"}).Decode(&chat)
// 	// if err != nil {
// 	// 	if err == mongo.ErrNoDocuments {
// 	_, err = db.tables["Chats"].InsertOne(ctx, ChatInfo{Messages: []api.Message{{MessageJSON: string([]byte(`{"message":"test message"}`))}},
// 		ChatID: "-1", Title: "Test chat", LastMessage: "test message", Participants: []string{}})
// 	if err != nil {
// 		log.Error.Println("createTestChat error:", err)
// 	}
// }

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

	// db.createTestChat() //Just for test

	// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	// _, _ = db.tables["Chats"].InsertOne(ctx, api,Chat{Messages: &api.Messages{Messages: []*api.Message{&api.Message{MessageJson: string([]byte(`{"message":"test message"}`))}}},
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

// func (db *DB) IsEmailUnique(email string) bool {
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	var result bson.M
// 	err := db.tables["Users"].FindOne(ctx, bson.M{"email": email}).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (db *DB) ValidateUser(email, password string) (bool, string) {
// 	log.Info.Println("Validate user", email, password)
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	var record bson.M
// 	if err := db.tables["Users"].FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&record); err != nil {
// 		log.Warning.Println(err)
// 		return false, ""
// 	}
// 	log.Info.Println(record)

// 	userID, ok := record["userid"].(string)
// 	if !ok {
// 		log.Warning.Println("ValidateUser error")
// 		return false, ""
// 	}

// 	return true, userID
// }

// func (db *DB) RegisterUser(data api.SignUp) (bool, string) {
// 	user := UserInfo{
// 		Email:    data.Email,
// 		Password: data.Password,
// 		Name:     data.Name,
// 		NickName: data.NickName,
// 		Photo:    data.Photo,
// 	}

// 	user.UserID = uuid.New().String()
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	_, err := db.tables["Users"].InsertOne(ctx, user)
// 	if err != nil {
// 		log.Warning.Println(err)
// 		return false, ""
// 	}

// 	return true, user.UserID
// }

// func (db *DB) AddUserToChat(userID string, chatID string) {
// 	var chat ChatInfo
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	err := db.tables["Chats"].FindOne(ctx, bson.M{"chatid": chatID}).Decode(&chat)
// 	if err != nil {
// 		log.Warning.Println("DB error:", err)
// 		return
// 	}

// 	chat.Participants = append(chat.Participants, userID)
// 	filter := bson.D{{"chatid", chatID}}
// 	update := bson.D{{"$set", bson.D{{"participants", chat.Participants}}}}
// 	_, err = db.tables["Chats"].UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Warning.Println("DB error:", err)
// 		return
// 	}

// 	var user api.UserInfo
// 	if err := db.tables["Users"].FindOne(ctx, bson.M{"userid": userID}).Decode(&user); err != nil {
// 		log.Warning.Println(err)
// 		return
// 	}

// 	user.Chats = append(user.Chats, chatID)
// 	filter = bson.D{{"userid", userID}}
// 	update = bson.D{{"$set", bson.D{{"chats", user.Chats}}}}
// 	_, err = db.tables["Users"].UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Warning.Println("DB error:", err)
// 		return
// 	}
// }

// func (db *DB) getChatInfo(chatID string) ChatInfo {
// 	log.Info.Println(chatID)
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	var chat ChatInfo
// 	err := db.tables["Chats"].FindOne(ctx, bson.M{"chatid": chatID}).Decode(&chat)
// 	if err != nil {
// 		log.Warning.Println("DB error:", err)
// 	}

// 	return chat
// }

// func (db *DB) GetUserChats(userID string) []api.Chat {
// 	log.Info.Println("GetUserChatsr", userID)
// 	result := make([]*api,Chat, 0)
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// 	var user api.UserInfo

// 	err := db.tables["Users"].FindOne(ctx, bson.M{"userid": userID}).Decode(&user)
// 	if err != nil {
// 		log.Warning.Println("DB error:", err)
// 	}

// 	for _, chatID := range user.Chats {
// 		chat := db.getChatInfo(chatID)
// 		result = append(result, &chat)
// 	}

// 	log.Info.Println(result)

// 	return api.Chats{Chats: result}
// }

// // If message_from is empty - return all available messages for chat
// // otherwise - messages from provided message_id
// func (db *DB) GetMessages(chatID string) []api.Message {
// 	log.Info.Printf("GetMessages %v < %v >\n", chatID)
// 	chat := db.getChatInfo(chatID)

// 	return chat.Messages
// }

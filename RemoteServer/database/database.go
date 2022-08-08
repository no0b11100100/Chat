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
	GetUserChats(string) []api.ChatInformation
	GetChatParticipants(string) []api.ParticipantInfo
	// GetUsers() []common.User
	// GetChatMessages(string) []common.Message
	// AddMessage(common.Message)
	// AddChat(common.Chat) error
	// AddUserToChat(string) error
	// LeaveChat(string, string) error
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

	userID, ok := record["user_id"].(string)
	if !ok {
		log.Warning.Println("ValidateUser error")
		return false, ""
	}

	return true, userID
}

func (db *DB) RegisterUser(user common.User) (bool, string) {
	user.ID = uuid.New().String()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.tables["Users"].InsertOne(ctx, user)
	if err != nil {
		log.Warning.Println(err)
		return false, ""
	}

	return true, user.ID
}

func (db *DB) GetUserChats(userID string) []api.ChatInformation {
	log.Info.Println("GetUserChatsr", userID)
	result := make([]api.ChatInformation, 0)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := db.tables["Users"].Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	for cur.Next(ctx) {
		var chat api.ChatInformation
		err = cur.Decode(&chat)
		if err != nil {
			log.Warning.Println("DB error:", err)
		}

		result = append(result, chat)
	}

	if err := cur.Err(); err != nil {
		log.Error.Println(err)
	}

	cur.Close(ctx)

	return result
}

func (db *DB) getParticipantInfo(userID string) api.ParticipantInfo {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var participant api.ParticipantInfo
	if err := db.tables["Users"].FindOne(ctx, bson.M{"user_id": userID}).Decode(&participant); err != nil {
		log.Warning.Println(err)
	}

	return participant
}

func (db *DB) GetChatParticipants(chatID string) []api.ParticipantInfo {
	log.Info.Println("GetChatParticipants", chatID)
	result := make([]api.ParticipantInfo, 0)
	var chat api.ChatInformation

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.tables["Chats"].FindOne(ctx, bson.M{"chat_id": chatID}).Decode(&chat)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	for _, userID := range chat.Participants() {
		result = append(result, db.getParticipantInfo(userID))
	}

	return result
}

package database

import (
	"Chat/RemoteServer/common"
	"context"
	"fmt"
	"log"
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
	// GetUserChats(string) []common.Chat
	// GetUsers() []common.User
	// GetChatMessages(string) []common.Message
	// GetChatParticipants(string) []common.Participant
	// AddMessage(common.Message)
	// AddChat(common.Chat) error
	// AddUserToChat(string) error
	// LeaveChat(string, string) error
}

func (db *DB) Connect() {
	fmt.Println("Connect database")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/db") //("mongodb://172.17.0.2:27017")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println("DB Connect error:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("DB Connect ping error:", err)
	}

	fmt.Println("Connected to MongoDB!")

	db.client = client

	users := client.Database("application").Collection("Users")
	chats := client.Database("application").Collection("Chats")

	db.tables["Users"] = users
	db.tables["Chats"] = chats

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func (db *DB) Close() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.client.Disconnect(ctx)

	if err != nil {
		log.Fatal("DB Close error:", err)
	}
	fmt.Println("Connection to MongoDB closed.")
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
	fmt.Println("Validate user", user)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var record bson.M
	if err := db.tables["Users"].FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&record); err != nil {
		fmt.Println(err)
		return false, ""
	}
	fmt.Println(record)

	userID, ok := record["user_id"].(string)
	if !ok {
		fmt.Println("ValidateUser error")
		return false, ""
	}

	return true, userID
}

func (db *DB) RegisterUser(user common.User) (bool, string) {
	user.ID = uuid.New().String()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.tables["Users"].InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	return true, user.ID
}

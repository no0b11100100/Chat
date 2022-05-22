package database

import (
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
	ValidateUser(User) (bool, string)
	RegisterUser(User) string
	GetUserChats(string) []Chat
	GetUsers() []User
	GetChatMessages(string) []Message
	GetChatParticipants(string) []Participant
	AddMessage(string, Message)
}

func (db *DB) Connect() {
	fmt.Println("Connect database")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
	err := db.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal("DB Close error:", err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func (db *DB) IsEmailUnique(email string) bool {
	err := db.tables["Users"].FindOne(context.TODO(), bson.M{"email": email})
	if err != nil {
		return false
	}
	return true
}
func (db *DB) ValidateUser(User) (bool, string) {
	return false, ""
}
func (db *DB) RegisterUser(User) string {
	uid := uuid.New()
	return uid.String()
}
func (db *DB) GetUserChats(string) []Chat {
	return nil
}
func (db *DB) GetUsers() []User {
	return nil
}
func (db *DB) GetChatMessages(string) []Message {
	return nil
}
func (db *DB) GetChatParticipants(string) []Participant {
	return nil
}
func (db *DB) AddMessage(string, Message) {}

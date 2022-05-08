package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

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
	ValidateUser(interface{}) bool
	RegisterUser(interface{})
	GetUserChats(string) []interface{}
	GetUsers() []interface{}
	GetChatMessages(string) []interface{}
	GetChatParticipants(string) []interface{}
	AddMessage(string, interface{})
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

func (db *DB) IsEmailUnique(string) bool {
	return false
}
func (db *DB) ValidateUser(interface{}) bool {
	return false
}
func (db *DB) RegisterUser(interface{}) {}
func (db *DB) GetUserChats(string) []interface{} {
	return nil
}
func (db *DB) GetUsers() []interface{} {
	return nil
}
func (db *DB) GetChatMessages(string) []interface{} {
	return nil
}
func (db *DB) GetChatParticipants(string) []interface{} {
	return nil
}
func (db *DB) AddMessage(string, interface{}) {}

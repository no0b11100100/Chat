package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	api "Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
)

// //https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

type DB struct {
	m                sync.RWMutex
	client           *mongo.Client
	userDatabase     *UserDatabase
	chatsDatabase    *ChatsDatabase
	calendarDatabase *CalendarDatabase
	todoListDataBase *TodolistDatabase
}

func NewDatabase() *DB {
	return &DB{
		m:                sync.RWMutex{},
		userDatabase:     NewUserDatabase(),
		chatsDatabase:    NewChatsDatabase(),
		calendarDatabase: NewCalendarDatabase(),
		todoListDataBase: NewTodolistDatabase(),
	}
}

type Database interface {
	interfaces.UserServiceDatabase
	interfaces.ChatServiceDatabase
	interfaces.CalendarServiceDatabase
	interfaces.TodoListServiceDataBase
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

	db.userDatabase.Connect(db.client, "application")
	db.chatsDatabase.Connect(db.client, "application")
	db.calendarDatabase.Connect(db.client, "application")
	db.todoListDataBase.Connect(db.client, "application")

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
	return db.userDatabase.IsEmailUnique(email)
}

func (db *DB) ValidateUser(email, password string) bool {
	return db.userDatabase.ValidateUser(email, password)
}

func (db *DB) RegisterUser(data api.SignUp) bool {
	return db.userDatabase.RegisterUser(data)
}

func (db *DB) AddUserToChat(userID string, chatID string) {
	db.chatsDatabase.AddUserToChat(userID, chatID)
	db.userDatabase.UpdateUserChats(userID, chatID)
}

func (db *DB) GetUserChats(userID string) []api.Chat {
	chats := db.userDatabase.GetUserChatIDs(userID)
	result := make([]api.Chat, 0)
	for _, chatID := range chats {
		result = append(result, db.chatsDatabase.GetChatInfo(chatID))
	}

	return result
}

// If message_from is empty - return all available messages for chat
// otherwise - messages from provided message_id
func (db *DB) GetMessages(chatID string) []api.Message {
	log.Info.Printf("GetMessages %v\n", chatID)
	chat := db.chatsDatabase.GetChatInfo(chatID)

	return chat.Messages
}

func (db *DB) AddMessage(msg api.Message) error {
	return db.chatsDatabase.AddMessage(msg)
}

func (db *DB) AddMeeting(userID string, meeting api.Meeting) error {
	fmt.Println("AddMeeting", userID)
	return db.calendarDatabase.AddMeeting(userID, meeting)
}

func (db *DB) GetMeetings(userID string, startDay string, endDay string) []api.Meeting {
	return db.calendarDatabase.GetMeetings(userID, startDay, endDay)
}

func (db *DB) AddList(userID string, list api.List) {
	db.todoListDataBase.AddList(userID, list)
}

func (db *DB) GetLists(userID string) []api.List {
	return db.todoListDataBase.GetLists(userID)
}

func (db *DB) AddTask(userID, listID string, task api.Task) {
	db.todoListDataBase.AddTask(userID, listID, task)
}

func (db *DB) GetListTasks(userID, listID string) []api.Task {
	return db.todoListDataBase.GetListTasks(userID, listID)
}

func (db *DB) SetTaskState(userID, listID, taskID string, state bool) {
	db.todoListDataBase.SetTaskState(userID, listID, taskID, state)
}

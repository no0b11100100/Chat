package database

import (
	api "Chat/Server/api"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodolistDatabase struct {
	*Base
}

func NewTodolistDatabase() *TodolistDatabase {
	return &TodolistDatabase{&Base{}}
}

func (db *TodolistDatabase) Connect(client *mongo.Client, database string) {
	db.Base.Connect(client, database, "Todolist")
}

func (db *TodolistDatabase) AddTask(userID, listID string, task api.Task) {}

func (db *TodolistDatabase) GetListTasks(userID, listID string) []api.Task {
	return make([]api.Task, 0)
}

func (db *TodolistDatabase) GetLists(userID string) []api.List {
	return make([]api.List, 0)
}

func (db *TodolistDatabase) AddList(userID string, list api.List) {}

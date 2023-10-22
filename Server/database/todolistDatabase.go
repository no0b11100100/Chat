package database

import (
	api "Chat/Server/api"
	log "Chat/Server/logger"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (db *TodolistDatabase) AddTask(userID, listID string, task api.Task) {
	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}
	fmt.Println("AddTask tag", userID, listID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	for _, list := range data.Lists {
		if list.Id == listID {
			list.Tasks = append(list.Tasks, task)
			break
		}
	}

	filter := bson.D{{"userID", userID}}
	update := bson.D{{"$set", bson.D{{"lists", data.Lists}}}}
	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

func (db *TodolistDatabase) GetListTasks(userID, listID string) []api.Task {
	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}
	fmt.Println("AddTask tag", userID, listID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return make([]api.Task, 0)
	}

	for _, list := range data.Lists {
		if list.Id == listID {
			return list.Tasks
		}
	}

	return make([]api.Task, 0)
}

func (db *TodolistDatabase) GetLists(userID string) []api.List {
	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}
	fmt.Println("GetLists tag", userID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return make([]api.List, 0)
	}
	log.Info.Println("Lists", data)
	return make([]api.List, 0)
}

func (db *TodolistDatabase) AddList(userID string, list api.List) {
	fmt.Println("AddList tag", userID, list.Title)

	filter := bson.D{{"userID", bson.M{"$exists": false}}}
	count, err := db.Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Warning.Println("Collection not exists")
		return
	}

	if count == 0 {
		document := bson.D{{"userID", userID}, {"lists", []api.List{list}}}
		_, err := db.Collection.InsertOne(context.Background(), document)
		if err != nil {
			log.Warning.Println("Can not add list")
			return
		}
		return
	}

	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = db.Collection.FindOne(ctx, bson.M{userID: userID}).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	data.Lists = append(data.Lists, list)
	filter = bson.D{{"userID", userID}}
	update := bson.D{{"$set", bson.D{{"lists", data.Lists}}}}
	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

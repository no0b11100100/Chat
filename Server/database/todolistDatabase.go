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

	for idx, list := range data.Lists {
		if list.Id == listID {
			data.Lists[idx].Tasks = append(data.Lists[idx].Tasks, task)
			break
		}
	}

	log.Info.Printf("Data %+v\n", data)

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
	fmt.Println("GetListTasks tag", userID, listID)
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
	return data.Lists
}

func (db *TodolistDatabase) AddList(userID string, list api.List) {
	fmt.Println("AddList tag", userID, list.Title)

	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}

	filter := bson.M{"userID": userID}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		if err == mongo.ErrNoDocuments {
			fmt.Println("Record not found")
			document := bson.D{{"userID", userID}, {"lists", []api.List{list}}}
			_, err = db.Collection.InsertOne(context.Background(), document)
			if err != nil {
				log.Warning.Println("Can not add list")
			}
		}
		return
	}

	data.Lists = append(data.Lists, list)
	updateFilter := bson.D{{"userID", userID}}
	update := bson.D{{"$set", bson.D{{"lists", data.Lists}}}}
	_, err = db.Collection.UpdateOne(ctx, updateFilter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

func (db *TodolistDatabase) SetTaskState(userID, listID, taskID string, state bool) {
	fmt.Println("SetTaskState", userID, listID, taskID, state)

	data := struct {
		UserID string     `bson:"userID"`
		Lists  []api.List `bson:"lists"`
	}{}
	fmt.Println("GetListTasks tag", userID, listID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.Collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&data)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}

	for listIdx, list := range data.Lists {
		if list.Id == listID {
			for taskIdx, task := range list.Tasks {
				if task.Id == taskID {
					data.Lists[listIdx].Tasks[taskIdx].IsCompleted = state
				}
			}
		}
	}

	updateFilter := bson.D{{"userID", userID}}
	update := bson.D{{"$set", bson.D{{"lists", data.Lists}}}}
	_, err = db.Collection.UpdateOne(ctx, updateFilter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

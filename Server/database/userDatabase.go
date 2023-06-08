package database

import (
	api "Chat/Server/api"
	log "Chat/Server/logger"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	*Base
}

func NewUserDatabase() *UserDatabase {
	return &UserDatabase{&Base{}}
}

func (db *UserDatabase) Connect(client *mongo.Client, database string) {
	db.Base.Connect(client, database, "Users")
}

func (db *UserDatabase) IsEmailUnique(email string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	userTagger := api.UserInfoTags()
	err := db.Collection.FindOne(ctx, bson.M{userTagger.Email: email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true
		}
	}
	return false
}

func (db *UserDatabase) ValidateUser(email, password string) bool {
	log.Info.Println("Validate user", email, password)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var record bson.M
	userTagger := api.UserInfoTags()
	if err := db.Collection.FindOne(ctx, bson.M{userTagger.Email: email, userTagger.Password: password}).Decode(&record); err != nil {
		log.Warning.Println(err)
		return false
	}
	log.Info.Println(record)
	return true
}

func (db *UserDatabase) RegisterUser(data api.SignUp) bool {
	user := api.UserInfo{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
		NickName: data.NickName,
		Photo:    data.Photo,
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Warning.Println(err)
		return false
	}

	return true
}

func (db *UserDatabase) UpdateUserChats(userID string, chatID string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user api.UserInfo
	userTagger := api.UserInfoTags()
	if err := db.Collection.FindOne(ctx, bson.M{userTagger.Email: userID}).Decode(&user); err != nil {
		log.Warning.Println(err)
		return
	}

	user.Chats = append(user.Chats, chatID)
	filter := bson.D{{userTagger.Email, userID}}
	update := bson.D{{"$set", bson.D{{userTagger.Chats, user.Chats}}}}
	_, err := db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Warning.Println("DB error:", err)
		return
	}
}

func (db *UserDatabase) GetUserChatIDs(userID string) []string {
	log.Info.Println("GetUserChatsr", userID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var user api.UserInfo
	userTagger := api.UserInfoTags()

	err := db.Collection.FindOne(ctx, bson.M{userTagger.Email: userID}).Decode(&user)
	if err != nil {
		log.Warning.Println("DB error:", err)
	}

	return user.Chats
}

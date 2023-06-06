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

type CalendarDatabase struct {
	*Base
}

func NewCalendarDatabase() *CalendarDatabase {
	return &CalendarDatabase{&Base{}}
}

func (db *CalendarDatabase) Connect(client *mongo.Client, database string) {
	db.Base.Connect(client, database, "Calendar")
}

/*
"series": [

	{
		"title": "name",
		"start": "date",
		"end": "" // if empty - active series, otherwise canceled
		participants: ["userID"]
	}

],
"meetings": []
*/

func (db *CalendarDatabase) AddMeeting(userID string, meeting api.Meeting) error {
	fmt.Println("AddMeeting", userID)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.Collection.InsertOne(ctx, meeting)
	if err != nil {
		log.Warning.Println(err)
		return err
	}
	return nil
}

func (db *CalendarDatabase) GetMeetings(userID string, startDay string, endDay string) []api.Meeting {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	meetings := db.getUserMeetings(ctx, userID, startDay, endDay)
	series := db.getUserSeriesMeetings(ctx, userID, startDay, endDay)

	return append(meetings, series...)
}

func (db *CalendarDatabase) getUserMeetings(ctx context.Context, userID string, startDay string, endDay string) []api.Meeting {
	tags := api.MeetingTags()
	cursor, err := db.Collection.Find(ctx, bson.M{tags.IsSeries: false})

	if err != nil {
		fmt.Println("getUserMeetings error", err)
		return make([]api.Meeting, 0)
	}

	defer cursor.Close(ctx)

	result := make([]api.Meeting, 0)

	for cursor.Next(ctx) {
		var m api.Meeting
		cursor.Decode(&m)
		if db.isUserInParticipants(userID, m.Participants) {
			result = append(result, m)
		}
	}

	log.Info.Println("getUserMeetings meetings count", len(result))

	return result
}

func (db *CalendarDatabase) getUserSeriesMeetings(ctx context.Context, userID string, startDay string, endDay string) []api.Meeting {
	tags := api.MeetingTags()
	cursor, err := db.Collection.Find(ctx, bson.M{tags.IsSeries: true})

	if err != nil {
		fmt.Println("getUserSeriesMeetings error", err)
		return make([]api.Meeting, 0)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		fmt.Println("getUserSeriesMeetings find meeting")
	}

	return make([]api.Meeting, 0)
}

func (db *CalendarDatabase) isUserInParticipants(userID string, participants []string) bool {
	for _, p := range participants {
		if p == userID {
			return true
		}
	}

	return false
}

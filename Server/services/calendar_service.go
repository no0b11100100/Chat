package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
	"fmt"
)

type CalendarService struct {
	database          interfaces.CalendarServiceDatabase
	serviceConnection *ServiceConnection
	notifiers         map[string]api.CalendarServiceNotifier
}

func NewCalendarService(serviceConnection *ServiceConnection, database interfaces.CalendarServiceDatabase) *CalendarService {
	s := &CalendarService{
		database:          database,
		serviceConnection: serviceConnection,
		notifiers:         make(map[string]api.CalendarServiceNotifier),
	}

	return s
}

func (calendar *CalendarService) CreateMeeting(_ api.ServerContext, m api.Meeting) api.ResponseStatus {
	for _, participant := range m.Participants {
		participantConnectID := calendar.serviceConnection.ConnectionIDByUserEmail.Request(participant)
		calendar.notifiers[participantConnectID].RecieveMeeting(m)
		if err := calendar.database.AddMeeting(participant, m); err != nil {
			fmt.Println("CreateMeeting error", participant, err)
		}
	}
	return api.OK
}

func (calendar *CalendarService) GetMeetings(_ api.ServerContext, userID string, daysRange api.DaysRange) []api.Meeting {
	log.Info.Println("GetMeetings", userID, daysRange.Start, daysRange.End)
	return calendar.database.GetMeetings(userID, daysRange.Start, daysRange.End)
}

func (calendar *CalendarService) HandleNewConnection(connectionID string, notifier api.CalendarServiceNotifier) {
	calendar.notifiers[connectionID] = notifier
}

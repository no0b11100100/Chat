package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
)

type CalendarService struct {
	database          interfaces.ChatServiceDatabase
	serviceConnection *ServiceConnection
	notifiers         map[string]api.CalendarServiceNotifier
}

func NewCalendarService(serviceConnection *ServiceConnection, database interfaces.ChatServiceDatabase) *CalendarService {
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
	}
	return api.OK
}

func (calendar *CalendarService) GetMeetings(api.ServerContext, string) []api.Meeting {
	return make([]api.Meeting, 0)
}

func (calendar *CalendarService) HandleNewConnection(connectionID string, notifier api.CalendarServiceNotifier) {
	calendar.notifiers[connectionID] = notifier
}

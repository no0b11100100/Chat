#pragma once

#include <memory>
#include "../communication/calendar_client.hpp"

class CalendarService final {
    std::unique_ptr<calendar::CalendarServiceStub> m_stub;
public:
    CalendarService(const std::string& addr)
    : m_stub{new calendar::CalendarServiceStub(addr)}
    {}

    ResponseStatus CreateMeeting(std::string name, std::vector<std::string> participants) {
        return m_stub->CreateMeeting(name, participants);
    }

    std::vector<calendar::Meeting> GetMeetings(std::string userID) {
        return m_stub->GetMeetings(userID);
    }

    void meetingUpdate(std::function<void(calendar::Meeting)> callback)
    {
        m_stub->SubscribeToRecieveMeetingEvent(callback);
    }

};

#pragma once

#include "common.hpp"

#include "types.hpp"

namespace calendar {

// enums

// structs
struct Meeting : public Types::ClassParser {
  std::string Title;
  std::vector<std::string> Participants;
  virtual json toJson() const override {
    json js({});
    js["title"] = Title;
    js["participants"] = Participants;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Title = std::string();
    else
      Title = static_cast<std::string>(js["title"]);
    if (js.isNull())
      Participants = std::vector<std::string>();
    else
      Participants = static_cast<std::vector<std::string>>(js["participants"]);
  }
};

class CalendarServiceStub : public Common::Base {
  std::vector<std::function<void(Meeting)>> m_RecieveMeetingCallbacks;

public:
  CalendarServiceStub() = default;

  CalendarServiceStub(const std::string &addr) : Common::Base(addr) {
    addSignalHandler("CalendarService.RecieveMeeting",
                     [this](json payload) { onRecieveMeeting(payload); });
  }

  void SubscribeToRecieveMeetingEvent(
      std::function<void(Meeting meeting)> callback) {
    m_RecieveMeetingCallbacks.push_back(callback);
  }

  ResponseStatus CreateMeeting(std::string name,
                               std::vector<std::string> participants) {
    Common::MessageData _message;
    _message.Payload = json::array({name, participants});
    _message.Endpoint = "CalendarService.CreateMeeting";
    return Request<ResponseStatus>(_message);
  }
  std::vector<Meeting> GetMeetings(std::string userID) {
    Common::MessageData _message;
    _message.Payload = json::array({userID});
    _message.Endpoint = "CalendarService.GetMeetings";
    return Request<std::vector<Meeting>>(_message);
  }

private:
  void onRecieveMeeting(json response) {
    int i = 0;
    Meeting meeting = response[i];
    ++i;

    for (const auto &call : m_RecieveMeetingCallbacks) {
      call(meeting);
    }
  }
};

}; // namespace calendar
#pragma once

#include "common.hpp"

#include "types.hpp"

namespace chat {

// enums

// structs
struct Message : public Types::ClassParser {
  std::string MessageJSON;
  std::string ChatID;
  std::string SenderID;
  virtual json toJson() const override {
    json js({});
    js["MessageJSON"] = MessageJSON;
    js["ChatID"] = ChatID;
    js["SenderID"] = SenderID;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      MessageJSON = std::string();
    else
      MessageJSON = static_cast<std::string>(js["MessageJSON"]);
    if (js.isNull())
      ChatID = std::string();
    else
      ChatID = static_cast<std::string>(js["ChatID"]);
    if (js.isNull())
      SenderID = std::string();
    else
      SenderID = static_cast<std::string>(js["SenderID"]);
  }
};

struct Chat : public Types::ClassParser {
  std::string ChatID;
  std::string Title;
  std::string SecondLine;
  std::string LastMessage;
  int UnreadedCount;
  std::string Cover;
  std::vector<std::string> Participants;
  std::vector<Message> Messages;
  virtual json toJson() const override {
    json js({});
    js["ChatID"] = ChatID;
    js["Title"] = Title;
    js["SecondLine"] = SecondLine;
    js["LastMessage"] = LastMessage;
    js["UnreadedCount"] = UnreadedCount;
    js["Cover"] = Cover;
    js["Participants"] = Participants;
    js["Messages"] = Messages;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      ChatID = std::string();
    else
      ChatID = static_cast<std::string>(js["ChatID"]);
    if (js.isNull())
      Title = std::string();
    else
      Title = static_cast<std::string>(js["Title"]);
    if (js.isNull())
      SecondLine = std::string();
    else
      SecondLine = static_cast<std::string>(js["SecondLine"]);
    if (js.isNull())
      LastMessage = std::string();
    else
      LastMessage = static_cast<std::string>(js["LastMessage"]);
    if (js.isNull())
      UnreadedCount = int();
    else
      UnreadedCount = static_cast<int>(js["UnreadedCount"]);
    if (js.isNull())
      Cover = std::string();
    else
      Cover = static_cast<std::string>(js["Cover"]);
    if (js.isNull())
      Participants = std::vector<std::string>();
    else
      Participants = static_cast<std::vector<std::string>>(js["Participants"]);
    if (js.isNull())
      Messages = std::vector<Message>();
    else
      Messages = static_cast<std::vector<Message>>(js["Messages"]);
  }
};

class ChatServiceStub : public Common::Base {
  std::vector<std::function<void(Message)>> m_RecieveMessageCallbacks;

public:
  ChatServiceStub() = default;

  ChatServiceStub(const std::string &addr) : Common::Base(addr) {
    addSignalHandler("ChatService.RecieveMessage",
                     [this](json payload) { onRecieveMessage(payload); });
  }

  void SubscribeToRecieveMessageEvent(
      std::function<void(Message message)> callback) {
    m_RecieveMessageCallbacks.push_back(callback);
  }

  ResponseStatus SendMessage(Message message) {
    Common::MessageData _message;
    _message.Payload = json::array({message});
    _message.Endpoint = "ChatService.SendMessage";
    return Request<ResponseStatus>(_message);
  }
  std::vector<Chat> GetUserChats(std::string userID) {
    Common::MessageData _message;
    _message.Payload = json::array({userID});
    _message.Endpoint = "ChatService.GetUserChats";
    return Request<std::vector<Chat>>(_message);
  }
  std::vector<Message> GetChatMessages(std::string chatID) {
    Common::MessageData _message;
    _message.Payload = json::array({chatID});
    _message.Endpoint = "ChatService.GetChatMessages";
    return Request<std::vector<Message>>(_message);
  }

private:
  void onRecieveMessage(json response) {
    int i = 0;
    Message message = response[i];
    ++i;

    for (const auto &call : m_RecieveMessageCallbacks) {
      call(message);
    }
  }
};

}; // namespace chat
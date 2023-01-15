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
  virtual json toJson() override {
    json js({});
    js["MessageJSON"] = MessageJSON;
    js["ChatID"] = ChatID;
    js["SenderID"] = SenderID;
    return js;
  }

  virtual void fromJson(json js) override {
    MessageJSON = static_cast<std::string>(js["MessageJSON"]);
    ChatID = static_cast<std::string>(js["ChatID"]);
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
  virtual json toJson() override {
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
    ChatID = static_cast<std::string>(js["ChatID"]);
    Title = static_cast<std::string>(js["Title"]);
    SecondLine = static_cast<std::string>(js["SecondLine"]);
    LastMessage = static_cast<std::string>(js["LastMessage"]);
    UnreadedCount = static_cast<int>(js["UnreadedCount"]);
    Cover = static_cast<std::string>(js["Cover"]);
    Participants = static_cast<std::vector<std::string>>(js["Participants"]);
    Messages = static_cast<std::vector<Message>>(js["Messages"]);
  }
};

class ChatServiceStub : public Common::Base {
  std::vector<std::function<void(Message)>> m_RecieveMessageCallbacks;

public:
  ChatServiceStub() = default;

  ChatServiceStub(const std::string &addr) : Common::Base(addr) {
    addSignalHandler("ChatService.RecieveMessage", [this](std::string payload) {
      onRecieveMessage(payload);
    });
  }

  void SubscribeToRecieveMessageEvent(
      std::function<void(Message message)> callback) {
    m_RecieveMessageCallbacks.push_back(callback);
  }

  ResponseStatus SendMessage(Message message) {
    json args = json::array({message});
    Common::MessageData _message;
    _message.Endpoint = "ChatService.SendMessage";
    _message.Payload = args.dump();
    return Request(_message).GetData<ResponseStatus>();
  }
  std::vector<Chat> GetUserChats(std::string userID) {
    json args = json::array({userID});
    Common::MessageData _message;
    _message.Endpoint = "ChatService.GetUserChats";
    _message.Payload = args.dump();
    return Request(_message).GetData<std::vector<Chat>>();
  }
  std::vector<Message> GetChatMessages(std::string chatID) {
    json args = json::array({chatID});
    Common::MessageData _message;
    _message.Endpoint = "ChatService.GetChatMessages";
    _message.Payload = args.dump();
    return Request(_message).GetData<std::vector<Message>>();
  }

private:
  void onRecieveMessage(const std::string &payload) {
    json response = json::parse(payload);
    int i = 0;
    Message message = response[i];
    ++i;

    for (const auto &call : m_RecieveMessageCallbacks) {
      call(message);
    }
  }
};

}; // namespace chat
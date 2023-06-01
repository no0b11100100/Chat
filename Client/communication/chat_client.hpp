#pragma once

#include "common.hpp"

#include "types.hpp"

namespace chat {

// enums
enum class CallStatus { Connected = 0, NotConnected = 1, Disconnected = 2 };

// structs
struct Timestamp : public Types::ClassParser {
  std::string Date;
  std::string Time;
  virtual json toJson() const override {
    json js({});
    js["date"] = Date;
    js["time"] = Time;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Date = std::string();
    else
      Date = static_cast<std::string>(js["date"]);
    if (js.isNull())
      Time = std::string();
    else
      Time = static_cast<std::string>(js["time"]);
  }
};

struct Message : public Types::ClassParser {
  json MessageJSON;
  std::string ChatID;
  std::string SenderID;
  Timestamp Date;
  virtual json toJson() const override {
    json js({});
    js["messagejson"] = MessageJSON;
    js["chatid"] = ChatID;
    js["senderid"] = SenderID;
    js["date"] = Date;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      MessageJSON = json();
    else
      MessageJSON = static_cast<json>(js["messagejson"]);
    if (js.isNull())
      ChatID = std::string();
    else
      ChatID = static_cast<std::string>(js["chatid"]);
    if (js.isNull())
      SenderID = std::string();
    else
      SenderID = static_cast<std::string>(js["senderid"]);
    if (js.isNull())
      Date = Timestamp();
    else
      Date = static_cast<Timestamp>(js["date"]);
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
    js["chatid"] = ChatID;
    js["title"] = Title;
    js["secondline"] = SecondLine;
    js["lastmessage"] = LastMessage;
    js["unreadedcount"] = UnreadedCount;
    js["cover"] = Cover;
    js["participants"] = Participants;
    js["messages"] = Messages;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      ChatID = std::string();
    else
      ChatID = static_cast<std::string>(js["chatid"]);
    if (js.isNull())
      Title = std::string();
    else
      Title = static_cast<std::string>(js["title"]);
    if (js.isNull())
      SecondLine = std::string();
    else
      SecondLine = static_cast<std::string>(js["secondline"]);
    if (js.isNull())
      LastMessage = std::string();
    else
      LastMessage = static_cast<std::string>(js["lastmessage"]);
    if (js.isNull())
      UnreadedCount = int();
    else
      UnreadedCount = static_cast<int>(js["unreadedcount"]);
    if (js.isNull())
      Cover = std::string();
    else
      Cover = static_cast<std::string>(js["cover"]);
    if (js.isNull())
      Participants = std::vector<std::string>();
    else
      Participants = static_cast<std::vector<std::string>>(js["participants"]);
    if (js.isNull())
      Messages = std::vector<Message>();
    else
      Messages = static_cast<std::vector<Message>>(js["messages"]);
  }
};

struct TextMessage : public Types::ClassParser {
  std::string Text;
  virtual json toJson() const override {
    json js({});
    js["text"] = Text;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Text = std::string();
    else
      Text = static_cast<std::string>(js["text"]);
  }
};

struct CallData : public Types::ClassParser {
  std::string Audio;
  virtual json toJson() const override {
    json js({});
    js["audio"] = Audio;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Audio = std::string();
    else
      Audio = static_cast<std::string>(js["audio"]);
  }
};

class ChatServiceStub : public Common::Base {
  std::vector<std::function<void(Message)>> m_RecieveMessageCallbacks;
  std::vector<std::function<void(CallData)>> m_NotifyCallDataCallbacks;
  std::vector<std::function<void()>> m_CallFromCallbacks;

public:
  ChatServiceStub() = default;

  ChatServiceStub(const std::string &addr) : Common::Base(addr) {
    addSignalHandler("ChatService.RecieveMessage",
                     [this](json payload) { onRecieveMessage(payload); });
    addSignalHandler("ChatService.NotifyCallData",
                     [this](json payload) { onNotifyCallData(payload); });
    addSignalHandler("ChatService.CallFrom",
                     [this](json payload) { onCallFrom(payload); });
  }

  void SubscribeToRecieveMessageEvent(
      std::function<void(Message message)> callback) {
    m_RecieveMessageCallbacks.push_back(callback);
  }
  void
  SubscribeToNotifyCallDataEvent(std::function<void(CallData data)> callback) {
    m_NotifyCallDataCallbacks.push_back(callback);
  }
  void SubscribeToCallFromEvent(std::function<void()> callback) {
    m_CallFromCallbacks.push_back(callback);
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
  CallStatus CallTo(std::string chatID, std::string callerID) {
    Common::MessageData _message;
    _message.Payload = json::array({chatID, callerID});
    _message.Endpoint = "ChatService.CallTo";
    return Request<CallStatus>(_message);
  }
  void SendCallData(CallData data) {
    Common::MessageData _message;
    _message.Payload = json::array({data});
    _message.Endpoint = "ChatService.SendCallData";
    Request(_message);
  }
  void HandleCallFrom(CallStatus status) {
    Common::MessageData _message;
    _message.Payload = json::array({status});
    _message.Endpoint = "ChatService.HandleCallFrom";
    Request(_message);
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
  void onNotifyCallData(json response) {
    int i = 0;
    CallData data = response[i];
    ++i;

    for (const auto &call : m_NotifyCallDataCallbacks) {
      call(data);
    }
  }
  void onCallFrom(json response) {
    int i = 0;

    for (const auto &call : m_CallFromCallbacks) {
      call();
    }
  }
};

}; // namespace chat
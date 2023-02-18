#pragma once

#include "defaultClient.hpp"
#include "json/Value/Value.h"

#include <algorithm>
#include <chrono>
#include <functional>
#include <future>
#include <iostream>
#include <memory>
#include <optional>
#include <string>
#include <unordered_map>
#include <vector>

namespace Common {
// enums
enum class MessageType { Request = 0, Notification = 1 };

// structs
struct MessageData : public Types::ClassParser {
  std::string ConnectionID;
  std::string Endpoint;
  std::string Topic;
  json Payload;
  MessageType Type;
  virtual json toJson() const override {
    json js({});
    js["connectionid"] = ConnectionID;
    js["endpoint"] = Endpoint;
    js["topic"] = Topic;
    js["payload"] = Payload;
    js["type"] = Type;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      ConnectionID = std::string();
    else
      ConnectionID = static_cast<std::string>(js["connectionid"]);
    if (js.isNull())
      Endpoint = std::string();
    else
      Endpoint = static_cast<std::string>(js["endpoint"]);
    if (js.isNull())
      Topic = std::string();
    else
      Topic = static_cast<std::string>(js["topic"]);
    if (js.isNull())
      Payload = json();
    else
      Payload = static_cast<json>(js["payload"]);
    if (js.isNull())
      Type = MessageType();
    else
      Type = static_cast<MessageType>(js["type"]);
  }
};

// client
class Base {
  std::unique_ptr<DefaultClient::TCPClient> m_connection;
  std::unordered_map<std::string, std::function<void(json)>> m_signals;
  std::unordered_map<std::string, std::function<void(json)>> m_waiters;

  inline static std::string CONNECTION_ID = "";
  static constexpr std::string_view BASE_SERVICE_ADDR = "localhost:1230";

public:
  Base(const std::string &addr)
      : m_connection{new DefaultClient::TCPClient(
            addr, [this](const std::string &s) { handleServerMessage(s); })} {
    if (Base::CONNECTION_ID.empty()) {
      qDebug() << "Init client"
               << QString::fromStdString(std::string(BASE_SERVICE_ADDR.data()));
      Base::CONNECTION_ID =
          m_connection->init(std::string(BASE_SERVICE_ADDR.data()));
      qDebug() << "CONNECTION_ID="
               << QString::fromStdString(Base::CONNECTION_ID);
    }
    m_connection->connectToServer();
  }

  Base() = default;

  template <class T> T Request(MessageData &message) {
    Request(message);
    std::promise<json> response;
    m_waiters[message.Topic] = [&response](json value) {
      response.set_value(value);
    };
    json response_paylaod = response.get_future().get();
    T result;
    result = static_cast<T>(response_paylaod);
    return result;
  }

  void Request(MessageData &message) {
    message.Type = MessageType::Request;
    message.ConnectionID = Base::CONNECTION_ID;
    message.Topic =
        std::to_string(std::chrono::duration_cast<std::chrono::milliseconds>(
                           std::chrono::system_clock::now().time_since_epoch())
                           .count());
    m_connection->sendPayload(message.toJson().dump());
  }

private:
  auto isSignalHandlable(const std::string &name) {
    auto it = std::find_if(
        m_signals.begin(), m_signals.end(),
        [&name](const std::pair<std::string, std::function<void(json)>> &p) {
          return p.first == name;
        });
    return it != m_signals.end()
               ? std::optional<std::function<void(json)>>{[this,
                                                           name](json payload) {
                   handle(name, payload);
                 }}
               : std::nullopt;
  }

  void handleServerMessage(const std::string &payload) {
    json js = json::parse(payload);
    MessageData data;
    data = js;
    if (data.Type == MessageType::Notification) {
      auto result = isSignalHandlable(data.Endpoint);
      if (result.has_value()) {
        std::cout << "Process notification topic " << data.Topic << std::endl;
        auto handler = result.value();
        handler(data.Payload);
      }
    } else {
      if (m_waiters.contains(data.Topic)) {
        std::cout << "Process topic " << data.Topic << std::endl;
        m_waiters[data.Topic](data.Payload);
        m_waiters.erase(data.Topic);
      }
    }
  }

  void handle(const std::string &name, json payload) {
    m_signals.at(name)(payload);
  }

protected:
  void addSignalHandler(std::string name, std::function<void(json)> handler) {
    m_signals.emplace(name, handler);
  }
};

}; // namespace Common
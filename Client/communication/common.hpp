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
  std::string Endpoint;
  std::string Topic;
  json Payload;
  MessageType Type;
  virtual json toJson() const override {
    json js({});
    js["Endpoint"] = Endpoint;
    js["Topic"] = Topic;
    js["Payload"] = Payload;
    js["Type"] = Type;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Endpoint = std::string();
    else
      Endpoint = static_cast<std::string>(js["Endpoint"]);
    if (js.isNull())
      Topic = std::string();
    else
      Topic = static_cast<std::string>(js["Topic"]);
    if (js.isNull())
      Payload = json();
    else
      Payload = static_cast<json>(js["Payload"]);
    if (js.isNull())
      Type = MessageType();
    else
      Type = static_cast<MessageType>(js["Type"]);
  }
};

// client
class Base {
  std::unique_ptr<DefaultClient::TCPClient> m_connection;
  std::unordered_map<std::string, std::function<void(json)>> m_signals;
  std::unordered_map<std::string, std::function<void(json)>> m_waiters;

public:
  Base(const std::string &addr)
      : m_connection{new DefaultClient::TCPClient(
            addr, [this](const std::string &s) { handleServerMessage(s); })} {
    m_connection->connectToServer();
  }

  Base() = default;

  template <class T> T Request(MessageData &message) {
    message.Type = MessageType::Request;
    message.Topic =
        std::to_string(std::chrono::duration_cast<std::chrono::milliseconds>(
                           std::chrono::system_clock::now().time_since_epoch())
                           .count());
    m_connection->sendPayload(message.toJson().dump());
    std::promise<json> response;
    m_waiters[message.Topic] = [&response](json value) {
      response.set_value(value);
    };
    json response_paylaod = response.get_future().get();
    T result;
    result = static_cast<T>(response_paylaod);
    return result;
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
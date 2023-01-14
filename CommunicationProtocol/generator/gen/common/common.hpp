#pragma once

#include "defaultClient.hpp"
#include "json/Value/Value.h"

#include <algorithm>
#include <chrono>
#include <functional>
#include <future>
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
  std::string Payload;
  MessageType Type;
  virtual json toJson() override {
    json js({});
    js["Endpoint"] = Endpoint;
    js["Topic"] = Topic;
    js["Payload"] = Payload;
    js["Type"] = Type;
    return js;
  }

  virtual void fromJson(json js) override {
    Endpoint = static_cast<std::string>(js["Endpoint"]);
    Topic = static_cast<std::string>(js["Topic"]);
    Payload = static_cast<std::string>(js["Payload"]);
    Type = static_cast<MessageType>(js["Type"]);
  }
};

// client
class Waiter {
  std::string m_topic;
  std::promise<std::string> m_response;

public:
  Waiter(std::string topicID) : m_topic{topicID} {}

  bool IsTopic(std::string topic) const { return m_topic == topic; }

  void SetResponse(std::string response) { m_response.set_value(response); }

  template <class T> T GetData() {
    auto future = m_response.get_future();
    auto response_paylaod = future.get();
    json response = json::parse(response_paylaod);
    T result = response;
    return result;
  }
};

class Base {
  std::unique_ptr<DefaultClient::TCPClient> m_connection;
  std::unordered_map<std::string, std::function<void(std::string)>> m_signals;
  std::vector<Waiter> m_waiters;

public:
  Base(const std::string &addr)
      : m_connection{new DefaultClient::TCPClient(
            [this](const std::string &s) { handleServerMessage(s); })} {
    m_connection->connectToServer();
  }

  Base() = default;

  Waiter &Request(MessageData &message) {
    message.Topic =
        std::to_string(std::chrono::duration_cast<std::chrono::milliseconds>(
                           std::chrono::system_clock::now().time_since_epoch())
                           .count());
    m_connection->sendPayload(message.toJson().dump());
    m_waiters.emplace_back(message.Topic);
    return m_waiters.back();
  }

private:
  auto isSignalHandlable(const std::string &name) {
    auto it = std::find_if(
        m_signals.begin(), m_signals.end(),
        [&name](
            const std::pair<std::string, std::function<void(std::string)>> &p) {
          return p.first == name;
        });
    return it != m_signals.end()
               ? std::optional<std::function<void(
                     std::string)>>{[this, name](const std::string &payload) {
                   handle(name, payload);
                 }}
               : std::nullopt;
  }

  void handleServerMessage(const std::string &payload) {
    json js = json::parse(payload);
    MessageData data;
    data = js;
    if (data.MessageType == Type::Notification) {
      auto result = isSignalHandlable(data.Endpoint);
      if (result.has_value()) {
        auto handler = result.value();
        handler(data.Payload);
      }
    } else {
      for (auto &waiter : m_waiters) {
        if (waiter.IsTopic(data.Topic)) {
          waiter.SetResponse(data.Payload);
          break;
        }
      }
    }
  }

  void handle(const std::string &name, const std::string &payload) {
    m_signals.at(name)(payload);
  }

protected:
  void addSignalHandler(std::string name,
                        std::function<void(std::string)> handler) {
    m_signals.emplace(name, handler);
  }
};

}; // namespace Common
#pragma once

#include <string>
#include <vector>
#include <functional>
#include <algothm>
#include <memory>
#include <future>
#include <unordered_map>

//enums
enum class Type
{
    Request = 0
    Notification = 1
};


//structs
struct Message : public ClassParser
{
    std::string Message;
virtual Value toJson() override
    {
        Value js({});
        js["Message"] = Message;
return js;
    }

    virtual void fromJson(Value js) override
    {
        Message = js["Message"];
}
};

struct Chat : public ClassParser
{
    std::string ChatId;
    std::vector<string> Participants;
virtual Value toJson() override
    {
        Value js({});
        js["ChatId"] = ChatId;
        js["Participants"] = Participants;
return js;
    }

    virtual void fromJson(Value js) override
    {
        ChatId = js["ChatId"];
        Participants = js["Participants"];
}
};

struct Status : public ClassParser
{
    int Status;
virtual Value toJson() override
    {
        Value js({});
        js["Status"] = Status;
return js;
    }

    virtual void fromJson(Value js) override
    {
        Status = js["Status"];
}
};

struct MessageData : public ClassParser
{
    std::string Endpoint;
    int Topic;
    std::string Payload;
    Type Type;
virtual Value toJson() override
    {
        Value js({});
        js["Endpoint"] = Endpoint;
        js["Topic"] = Topic;
        js["Payload"] = Payload;
        js["Type"] = Type;
return js;
    }

    virtual void fromJson(Value js) override
    {
        Endpoint = js["Endpoint"];
        Topic = js["Topic"];
        Payload = js["Payload"];
        Type = js["Type"];
}
};


//client
class Waiter
{
    int m_topic;
    std::promise<std::string> m_response;
public:
    Waiter(int topicID)
    : m_topic{topicID}
    {}

    bool IsTopic(int topic) const { return m_topic == topic; }

    void SetResponse(std::string response)
    {
        m_response.set_value(response);
    }

    template<class T>
    T GetData()
    {
        auto future = m_response.get_future();
        auto response_paylaod = future.get();
        json response = json::parse(response_paylaod);
        T result;
        result = response;
        return result;
    }

    //TODO
    template<void>
    void GetData()
    {}
};


class Base
{
    std::unique_ptr<QTcpSocket> m_connection;
    std::unordered_map<std::string, std::function<void(std::string)>> m_signals;
public:
    Base(const std::string& addr)
    {
        //TODO: create connection
        std::thread([this](){recieve();}).detach();
    }

    const Waiter& Request(payload std::string)
    {
        int requestID = 0;
        m_connection->send(paylaod);
        w = Waiter(requestID);
        waiters.push_back(w);
        return waiters.back();
    }

private:
    void recieve()
    {
        while(m_connection->read()){
            //TODO: unmarshal to struct DataStruct
            MessageData data;
            if (data.type == Notification) {
                auto result = isSignalHandlable(data.Endpoint);
                if(result.has_value())
                {
                    auto handler = result.value();
                    handler(data.Payload);
                }
            }
            else
            {
                for(auto& waiter : waiters)
                {
                    if waiter.IsTopic(data.Topic) {
                        waiter.SetResponse(data.Payload);
                        break;
                    }
                }
            }
        }
    }

    std::optional<std::function<void(std::string)>> isSignalHandlable(const std::string& name)
    {
        auto it = std::find_if(m_signals.begin(), m_signals.end(), [&name](const std::pair<std::string, std::function<void(std::string)>>& p)
        {
            return p.first == name;
        });
        if (it != m_signals.end())
            return std::optional<std::function<void(std::string)>>{[this, name](const std::string& payload){ handle(name, payload); }};
        return std::nullopt;
    }

    void handle(const std::string& name, const std::string& payload)
    {
        m_signals.at(name)(paylaod);
    }

protected:
    void addSignalHandler(std::string name, std::function<void(std::string)> handler)
    {
        m_signals.emplace(name, handler);
    }
};


class ChatServiceStub : Base
{
    std::vector<std::function<void(Message message,std::string data,)>> m_NotifyMessageCallbacks;
public:
    ChatService(const std::string& addr)
        : Base(addr)
    {
        addSignalHandler("ChatService.NotifyMessage", [this](std::string payload){ onNotifyMessage(payload); })
    }

    void SubscribeToNotifyMessageEvent(std::function<void(Message message,std::string data,)> callback)
    {
        m_NotifyMessageCallbacks.push_back(callback);
    }

    Status SendMessage(message Message,)
    {
        json args = json::array({message Message,});
        MessageData message;
        message.Endpoint = "ChatService.SendMessage";
        message.Topic = 0;//TODO
        message.Payload = args.dump();
        return Request(message.toJson().dump()).GetData<Status>();
    }
    std::vector<Chat> GetUserChats(userID string,value int,)
    {
        json args = json::array({userID string,value int,});
        MessageData message;
        message.Endpoint = "ChatService.GetUserChats";
        message.Topic = 0;//TODO
        message.Payload = args.dump();
        return Request(message.toJson().dump()).GetData<std::vector<Chat>>();
    }

private:
    void onNotifyMessage(const std::string& payload)
    {
        json response = json::parse(payload);
        int i = 0;
        Message message = response[i];
        ++i;
        std::string data = response[i];
        ++i;

        for(const auto& call : m_NotifyMessageCallbacks)
        {
            call(Message message,std::string data,);
        }
    }
};

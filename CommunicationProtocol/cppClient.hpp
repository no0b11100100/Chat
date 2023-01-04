#pragma once

class Waiter
{
    int m_topic;
    std::promise<std::string> m_response;
public:
    Waiter(int topicID)
    : m_topic{topicID}
    {}

    bool IsTopic(int topic) const { return m_topic == topic; }

    void SetResponse(string response)
    {
        m_response.set_value(response);
    }

    template<class T>
    T GetData()
    {
        auto future = m_response.get_future();
        auto response_paylaod = future.get();
        //unmarshal json to struct
        return T();
    }

    template<void>
    void GetData()
    {}
};

class Base
{
    std::unique_ptr<QTcpSocket> m_connection;
    std::unordered_map<std::string, std::function<void(std::string)>> m_signals;
public:
    Base(const string& addr)
    {
        //TODO: create connection
        std::thread([this](){recieve();}).detach();
    }

    const Waiter& Request(payload string)
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
            //unmarshal to struct DataStruct
            /*
                struct DataStruct {
                    topic int
                    response string
                    name string
                    type enum -> Response, notification
                }
            */
            DataStruct data;
            if (data.type == Notification) {
                auto result = isSignalHandlable(data.name);
                //TODO: call handler
            }
            else
            {
                for(auto& waiter : waiters)
                {
                    if waiter.IsTopic(data.topic) {
                        waiter.SetResponse(data.response);
                    }
                }
            }
        }
    }

    std::optional<std::function<void(std::string)>> isSignalHandlable(const std::string& name)
    {
        // check if signal in list
        auto fn = [this, name](const std::string& payload){ handle(name, payload); };
    }

    void handle(const std::string& name, const std::string& payload)
    {
        m_signals.at(name)(paylaod);
    }

protected:
    void addSignalHandler(std::string name, std::function<void(std::string)> handler)
    {}
};


class ChatServiceStub : Base
{
    std::vector<std::function<void(Message message)>> m_notifyMessageCallbacks;
public:
    ChatService(const string& addr)
        : Base(addr)
    {
        addSignalHandler("ChatService.notifyMessage", [this](std::string payload){ onNotifyMessage(payload); })
    }

    void onNotifyMessage(std::string payload)
    {
        //unmarshal payload to args
        // recieve data from server
        for(const auto& call : m_notifyMessageCallbacks)
        {
            call(<args>);
        }
    }

private:
    void sendMessage(Message message)
    {
        Request(message.dump()).GetData<void>();
    }

    std::vector<Chat> getUserChats(userID string)
    {
        auto response = Request(userID).GetData<std::vector<Chat>>();
        response.push_back(Chat());
        return response;
    }

    void SubscribeToNotifyMessageEvent(std::function<void(Message message)> callback)
    {
        m_notifyMessageCallbacks.push_back(callback);
    }
};

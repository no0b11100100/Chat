#pragma once

#include <memory>
#include "../communication/chat_client.hpp"

class ChatService final {
    std::unique_ptr<chat::ChatServiceStub> m_stub;
public:
    ChatService(const std::string& addr)
    : m_stub{new chat::ChatServiceStub(addr)}
    {}

    ResponseStatus sendMessage(std::string chatID, std::string sender, json message, chat::Timestamp timestamp) {
        chat::Message msg;
        msg.MessageJSON = message;
        msg.ChatID = chatID;
        msg.SenderID = sender;
        msg.Date = timestamp;
        return m_stub->SendMessage(msg);
    }

    std::vector<chat::Chat> getUserChats(std::string userID) {
        return m_stub->GetUserChats(userID);
    }

    std::vector<chat::Message> getChatMessages(std::string chatID) {
        return m_stub->GetChatMessages(chatID);
    }

    void messageUpdate(std::function<void(chat::Message)> callback)
    {
        m_stub->SubscribeToRecieveMessageEvent(callback);
    }

    chat::CallStatus CallTo(std::string chatID, std::string callerID) {
        return chat::CallStatus::Connected; //m_stub->CallTo(chatID, callerID);
    }

    void SendCallData(chat::CallData data) {
        m_stub->SendCallData(data);
    }

    void handleNotifyCallData(std::function<void(chat::CallData)> callback) {
        m_stub->SubscribeToNotifyCallDataEvent(callback);
    }
};

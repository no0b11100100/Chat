#pragma once

#include "ChatModel.hpp"
#include "ChatListModel.hpp"
#include "NotificationModel/NotificationModel.hpp"
#include "../../../services/Client.hpp"

#include <chrono>
#include <thread>

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel CONSTANT)
    Q_PROPERTY(QObject* chatListModel READ chatList CONSTANT)
    Q_PROPERTY(QObject* notificationListModel READ notificationListModel CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(Client* client, QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(parent)},
          m_chatList{new ChatListModel(parent)},
          m_notificationModel{new NotificationModel(parent)},
          m_client{client}
    {
        QObject::connect(m_chatList.get(), &ChatListModel::chatSelected, this, &BaseScreen::setChat);
        QObject::connect(m_chatModel.get(), &ChatModel::sendingMessage, this, &BaseScreen::sendMessage);
        m_client->chatService().messageUpdate([this](chat::Message msg){handleMessageNotification(msg);});
        // std::thread([this](){
        //         m_client->chatService().MessagesUpdated([this](chat::ExchangedMessage msg){handleMessageNotification(msg);});
        //     }).detach();
    }

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }
    QObject* notificationListModel() { return m_notificationModel.get(); }

    QString name() const { return "BaseScreen"; }

    void SetUser(user::Response userData)
    {
        // std::this_thread::sleep_for(std::chrono::seconds(20));
        auto chats = m_client->chatService().getUserChats(userData.Info.UserID);
        m_chatList->SetChats(chats);
    }


public slots:
    void setChat(const Header& header, QString chatID)
    {
        m_chatModel->SaveSelectedChatID(chatID);
        m_chatModel->SetHeader(header);
        auto messages = m_client->chatService().getChatMessages(chatID.toStdString());//, "", chat::Direction::Forward);
        m_chatModel->SetMessages(messages);
    }

    void sendMessage(QString chatID, QString message)
    {
        m_chatList->SetLastMessage(chatID, message);
        //TODO: add message properly
        //TODO: move logic for Message to ChatModel
        // chat::Message msg;
        std::string message_json = "{\"message\":\"" + message.toStdString() + "\"}";
        // mfg.MessageJSON = message_json;
        // msg.set_message_json(message_json);
        qDebug() << "Before SendMessage";
        m_client->chatService().sendMessage(chatID.toStdString(), message_json);
    }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
    std::unique_ptr<NotificationModel> m_notificationModel;
    Client* m_client;

    void handleMessageNotification(chat::Message message)
    {
        qDebug() << "handleMessage";
        m_chatModel->AddMessage(message);
        auto s = message.MessageJSON;
        QJsonDocument object = QJsonDocument::fromJson(QByteArray(s.data(), int(s.size())));
        QJsonObject message_json = object.object();
        std::string text = message_json["message"].toString().toStdString();
        m_chatList->SetLastMessage(QString::fromStdString(message.ChatID), QString::fromStdString(text));
    }
};

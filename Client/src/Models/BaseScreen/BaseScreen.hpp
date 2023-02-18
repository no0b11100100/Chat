#pragma once

#include "ChatModel.hpp"
#include "ChatListModel.hpp"
#include "NotificationModel/NotificationModel.hpp"
#include "CalendarModel/CalendarModel.hpp"
#include "../../../services/Client.hpp"

#include <chrono>
#include <thread>

class UserInfo {
    user::UserInfo m_info;
public:
    UserInfo(user::UserInfo info)
    : m_info{info}
    {}

    UserInfo() = default;

    user::UserInfo Info() const { return m_info; }

};

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel CONSTANT)
    Q_PROPERTY(QObject* chatListModel READ chatList CONSTANT)
    Q_PROPERTY(QObject* notificationListModel READ notificationListModel CONSTANT)
    Q_PROPERTY(QObject* calendarModel READ calendarModel CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(Client* client, QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(client->chatService(), parent)},
          m_chatList{new ChatListModel(parent)},
          m_notificationModel{new NotificationModel(parent)},
          m_calendarModel{new CalendarModel(client->calendarService(), parent)},
          m_client{client}
    {
        QObject::connect(m_chatList.get(), &ChatListModel::chatSelected, this, &BaseScreen::setChat);
        QObject::connect(m_chatModel.get(), &ChatModel::sendingMessage, this, &BaseScreen::sendMessage);
        QObject::connect(this, &BaseScreen::receiveMessage, this, &BaseScreen::handleMessageNotification);
        m_client->chatService().messageUpdate([this](chat::Message msg){emit receiveMessage(msg);});
    }

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }
    QObject* notificationListModel() { return m_notificationModel.get(); }
    QObject* calendarModel() { return m_calendarModel.get(); }

    QString name() const { return "BaseScreen"; }

    void SetUser(user::UserInfo userData)
    {
        // std::this_thread::sleep_for(std::chrono::seconds(20));
        auto chats = m_client->chatService().getUserChats(userData.UserID);
        m_chatList->SetChats(chats);
    }

signals:
    void receiveMessage(chat::Message);

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
        chat::TextMessage message_json;
        message_json.Text = message.toStdString();
        qDebug() << "Before SendMessage";
        m_client->chatService().sendMessage(chatID.toStdString(), message_json.toJson());
    }

    void handleMessageNotification(chat::Message message)
    {
        qDebug() << "handleMessage";
        m_chatModel->AddMessage(message);
        json js = message.MessageJSON;
        chat::TextMessage textMessage;
        textMessage = js;
        m_chatList->SetLastMessage(QString::fromStdString(message.ChatID), QString::fromStdString(textMessage.Text));
    }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
    std::unique_ptr<NotificationModel> m_notificationModel;
    std::unique_ptr<CalendarModel> m_calendarModel;
    Client* m_client;
};

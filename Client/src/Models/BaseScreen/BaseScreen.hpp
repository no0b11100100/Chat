#pragma once

#include "ChatModel/ChatModel.hpp"
#include "ChatListModel/ChatListModel.hpp"
#include "NotificationModel/NotificationModel.hpp"
#include "CalendarModel/CalendarModel.hpp"
#include "TodoListModel/TodoListModel.hpp"
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
    Q_PROPERTY(QObject* todoListModel READ todoListModel CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(Client* client, QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(client->chatService(), parent)},
          m_chatList{new ChatListModel(client->chatService(), parent)},
          m_notificationModel{new NotificationModel(parent)},
          m_calendarModel{new CalendarModel(client->calendarService(), parent)},
          m_todoListModel{new TodoListModel(client->todoListService(), parent)}
    {
        QObject::connect(m_chatList.get(), &ChatListModel::chatSelected, m_chatModel.get(), &ChatModel::setChat);
        QObject::connect(m_chatModel.get(), &ChatModel::sendingMessage, m_chatList.get(), &ChatListModel::updateLastMessage);
    }

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }
    QObject* notificationListModel() { return m_notificationModel.get(); }
    QObject* calendarModel() { return m_calendarModel.get(); }
    QObject* todoListModel() { return m_todoListModel.get(); }

    QString name() const { return "BaseScreen"; }

    void SetUser(user::UserInfo userData)
    {
        // std::this_thread::sleep_for(std::chrono::seconds(20));
        qDebug() << "SetUser" << QString::fromStdString(userData.UserID);
        m_chatList->SetChats(userData.UserID);
        m_chatModel->SetUserID(userData.UserID);
        m_todoListModel->SetLists(userData.UserID);
    }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
    std::unique_ptr<NotificationModel> m_notificationModel;
    std::unique_ptr<CalendarModel> m_calendarModel;
    std::unique_ptr<TodoListModel> m_todoListModel;
};

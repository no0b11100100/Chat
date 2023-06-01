#pragma once

#include <QObject>
#include "../../../communication/chat_client.hpp"
#include "../utils/utils.hpp"

class ChatInformation : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString lastMessage READ lastMessage NOTIFY lastMessageChanged)
    Q_PROPERTY(QString lastMessageTime READ lastMessageTime NOTIFY lastMessageChanged)
    Q_PROPERTY(QString id READ id CONSTANT)


public:
    ChatInformation(const QString& id, const QString& title, const chat::LastChatMessage& lastMessage, QObject* parent = nullptr)
        : QObject{parent},
        m_title{title},
        m_id{id},
        m_lastMessage{lastMessage}
    {}

    QString title() const { return m_title; }
    QString id() const { return m_id; }
    QString lastMessage() const { return QString::fromStdString(m_lastMessage.Message); }
    QString lastMessageTime() const { return convertTime(m_lastMessage.Date.Time); }

    void UpdateLastMessage(chat::LastChatMessage message)
    {
        m_lastMessage = message;
        emit lastMessageChanged();
    }

signals:
    void lastMessageChanged();

private:
    QString m_title;
    QString m_id;
    chat::LastChatMessage m_lastMessage;
};

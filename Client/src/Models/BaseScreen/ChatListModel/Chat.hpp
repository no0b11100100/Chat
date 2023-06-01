#pragma once

#include <QObject>

class ChatInformation : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString lastMessage READ lastMessage NOTIFY lastMessageChanged)
    Q_PROPERTY(QString lastMessageTime READ lastMessageTime NOTIFY lastMessageChanged)
    Q_PROPERTY(QString id READ id CONSTANT)


public:
    ChatInformation(const QString& id, const QString& title, const QString& lastMessage, QObject* parent = nullptr)
        : QObject{parent},
        m_title{title},
        m_id{id},
        m_lastMessage{lastMessage}
    {}

    QString title() const { return m_title; }
    QString id() const { return m_id; }
    QString lastMessage() const { return m_lastMessage; }
    QString lastMessageTime() const { return m_lastMessageTime; }

    void UpdateLastMessage(QString message, QString time)
    {
        m_lastMessage = message;
        m_lastMessageTime = time;
        emit lastMessageChanged();
    }

signals:
    void lastMessageChanged();

private:
    QString m_title;
    QString m_id;
    QString m_lastMessage;
    QString m_lastMessageTime;
};

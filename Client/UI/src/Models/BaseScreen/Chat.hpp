#pragma once

#include <QObject>

class ChatInfo : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString lastMessage READ lastMessage NOTIFY lastMessageChanged)
    Q_PROPERTY(QString id READ id CONSTANT)

public:
    ChatInfo(const QString& id, const QString& title, QObject* parent = nullptr)
        : QObject{parent},
        m_title{title},
        m_id{id}
    {}

    QString title() const { return m_title; }
    QString id() const { return m_id; }
    QString lastMessage() const { return m_lastMessage; }

    void UpdateLastMessage(QString message)
    {
        m_lastMessage = message;
        emit lastMessageChanged();
    }

signals:
    void lastMessageChanged();

private:
    QString m_title;
    QString m_id;
    QString m_lastMessage;
};

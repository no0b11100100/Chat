#pragma once

#include <QObject>

class Notification : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString text READ text CONSTANT)
    Q_PROPERTY(QString senderName READ senderName CONSTANT)
public:
    Notification(QString text, QString senderName, QObject* parent = nullptr)
     : QObject{parent},
     m_text{text},
     m_sender{senderName}
    {}

    QString text() const { return m_text; }
    QString senderName() const { return m_sender; }

private:
    QString m_text;
    QString m_sender;
    QString m_senderAvatar;
    QString m_time;
};

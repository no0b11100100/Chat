#pragma once

#include <QObject>

class SimpleMessage : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString message READ message CONSTANT)
    Q_PROPERTY(bool sendByMe READ sendByMe CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
    Q_PROPERTY(QString time READ time CONSTANT)

public:
    SimpleMessage(const QString& message, bool sendByMe, QString time, QObject* parent = nullptr)
        : QObject{parent},
        m_message{message},
        m_sendByMe{sendByMe},
        m_time{time}
    {}

    QString message() const { return m_message; }
    QString name() const { return "SimpleMessage"; }
    QString time() const { return m_time; }
    bool sendByMe() const { return m_sendByMe; }

private:
    QString m_message;
    bool m_sendByMe;
    QString m_time;
};
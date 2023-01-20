#pragma once

#include <QObject>

class SimpleMessage : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString message READ message CONSTANT)
    Q_PROPERTY(bool sendByMe READ sendByMe CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    SimpleMessage(const QString& message, bool sendByMe, QObject* parent = nullptr)
        : QObject{parent},
        m_message{message},
        m_sendByMe{sendByMe}
    {}

    QString message() const { return m_message; }
    QString name() const { return "SimpleMessage"; }
    bool sendByMe() const { return m_sendByMe; }

private:
    QString m_message;
    bool m_sendByMe;
};
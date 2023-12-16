#pragma once

#include <QObject>

class DateMessage : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString message READ message CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)

public:
    DateMessage(const QString& message, QObject* parent = nullptr)
        : QObject{parent},
        m_message{message}
    {}

    QString message() const { return m_message; }
    QString name() const { return "DateMessage"; }

private:
    QString m_message;
};

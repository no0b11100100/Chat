#pragma once

#include <QObject>

class ChatInfo : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)

public:
    ChatInfo(const QString& title, QObject* parent = nullptr)
        : QObject{parent},
        m_title{title}
    {}

    QString title() const { return m_title; }

private:
    QString m_title;
};
#pragma once

#include <QObject>
#include "../../../../communication/todolist_client.hpp"

class Task : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString id READ id CONSTANT)

public:
    Task(QString title, QString id, QObject* parent = nullptr)
    : QObject{parent},
    m_title{title},
    m_id{id}
    {}

    QString title() const { return m_title; }
    QString id() const { return m_id; }

private:
    QString m_title;
    QString m_id;

};

#pragma once

#include <QObject>
#include "../../../../communication/todolist_client.hpp"

class Task : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)

public:
    Task(QString title, QObject* parent = nullptr)
    : QObject{parent},
    m_title{title}
    {}

    QString title() const { return m_title; }

private:
    QString m_title;

};

#pragma once

#include <QObject>
#include "../../../../communication/todolist_client.hpp"

class List : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString id READ id CONSTANT)
public:
    List(todolist::List list, QObject* parent = nullptr)
    : QObject{parent},
    m_list{list}
    {}

    QString id() const { return QString::fromStdString(m_list.Id); }
    QString title() const { return QString::fromStdString(m_list.Title); }

private:
    todolist::List m_list;
};
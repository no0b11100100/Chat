#pragma once

#include <QObject>
#include "../../../../communication/todolist_client.hpp"

class Task : public QObject
{
    Q_OBJECT

public:
    Task(QObject* parent = nullptr)
    : QObject{parent}
    {}

};
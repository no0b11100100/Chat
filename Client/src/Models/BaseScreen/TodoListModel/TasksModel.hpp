#pragma once

#include <QObject>

class TasksModel  : public QObject
{
    Q_OBJECT
public:
    TasksModel(QObject* parent = nullptr)
    : QObject{parent}
    {}
};

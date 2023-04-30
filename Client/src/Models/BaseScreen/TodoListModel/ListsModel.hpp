#pragma once

#include <QObject>

class ListsModel  : public QObject
{
    Q_OBJECT
public:
    ListsModel(QObject* parent = nullptr)
    : QObject{parent}
    {}
};

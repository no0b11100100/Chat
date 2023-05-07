#pragma once

#include <QObject>
#include <QAbstractListModel>

#include <vector>
#include <memory>

#include "../../../../services/TodoListService.hpp"
#include "Task.hpp"

class TasksModel : public QAbstractListModel
{
    Q_OBJECT
public:
    TasksModel(TodoListService& client, QObject* parent = nullptr)
    : QAbstractListModel{parent},
    m_client{client}
    {}

    int rowCount(const QModelIndex& parent) const override
    {
        return m_tasks.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        qDebug() << role;
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue(m_tasks.at(index.row()).get());
    }

public slots:
    void setListTasks(QString listID)
    {
        qDebug() << "Request tasks for list" << listID;
        auto listTasks = m_client.GetTasks(listID.toStdString());
    }

private:
    TodoListService& m_client;
    std::vector<std::unique_ptr<Task>> m_tasks;
};

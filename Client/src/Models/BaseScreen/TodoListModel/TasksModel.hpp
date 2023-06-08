#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <QDebug>

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

    Q_INVOKABLE void addTask(QString title)
    {
        qDebug() << "Add task" << title;
        emit beginResetModel();
        m_tasks.emplace_back(std::make_unique<Task>(title));
        emit endResetModel();
        todolist::Task task;
        task.Description = title.toStdString();
        emit addedTask(task);
    }

    Q_INVOKABLE void setTaskState(QString taskID, bool state) {
        qDebug() << "changeTaskState" << taskID << state;
        emit changeTaskState(taskID, state);
    }

signals:
    void addedTask(todolist::Task);
    void changeTaskState(QString, bool);

public slots:
    void setListTasks(QString userID, QString listID)
    {
        qDebug() << "Request tasks for list" << listID;
        auto listTasks = m_client.GetTasks(userID.toStdString(), listID.toStdString());
        m_tasks.clear();
        emit beginResetModel();
        for(const auto& task : listTasks)
        {
            m_tasks.emplace_back(std::make_unique<Task>(QString::fromStdString(task.Description)));
        }
        emit endResetModel();
    }

private:
    TodoListService& m_client;
    std::vector<std::unique_ptr<Task>> m_tasks;
};

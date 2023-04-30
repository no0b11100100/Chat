#pragma once

#include <QObject>

#include "../../../../services/TodoListService.hpp"

#include "ListsModel.hpp"
#include "TasksModel.hpp"

class TodoListModel : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* lists READ lists CONSTANT)
    Q_PROPERTY(QObject* tasks READ tasks CONSTANT)
public:
    TodoListModel(TodoListService& client, QObject* parent = nullptr)
    : QObject{parent},
    m_client{client},
    m_tasksModel{new TasksModel(this)},
    m_listsModel{new ListsModel(this)}
    {}

    QObject* lists() { return m_listsModel.get(); }
    QObject* tasks() { return m_tasksModel.get(); }

private:
    TodoListService& m_client;
    std::unique_ptr<TasksModel> m_tasksModel;
    std::unique_ptr<ListsModel> m_listsModel;
};

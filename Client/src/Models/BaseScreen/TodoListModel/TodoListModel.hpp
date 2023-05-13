#pragma once

#include <QObject>

#include "ListsModel.hpp"
#include "TasksModel.hpp"

class TodoListModel : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* listsModel READ listsModel CONSTANT)
    Q_PROPERTY(QObject* tasksModel READ tasksModel CONSTANT)
public:
    TodoListModel(TodoListService& client, QObject* parent = nullptr)
    : QObject{parent},
    m_tasksModel{new TasksModel(client, this)},
    m_listsModel{new ListsModel(client, this)}
    {
        QObject::connect(m_listsModel.get(), &ListsModel::listSelected, m_tasksModel.get(), &TasksModel::setListTasks);
        QObject::connect(m_tasksModel.get(), &TasksModel::addedTask, m_listsModel.get(), &ListsModel::addNewTask);
    }

    QObject* listsModel() { return m_listsModel.get(); }
    QObject* tasksModel() { return m_tasksModel.get(); }

    void SetLists(std::string userID)
    {
        m_listsModel->SetLists(userID);
    }

private:
    std::unique_ptr<TasksModel> m_tasksModel;
    std::unique_ptr<ListsModel> m_listsModel;
};

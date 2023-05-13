#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <QDebug>

#include <vector>
#include <memory>

#include "../../../../services/TodoListService.hpp"
#include "List.hpp"

class ListsModel : public QAbstractListModel
{
    Q_OBJECT
public:
    ListsModel(TodoListService& client, QObject* parent = nullptr)
    : QAbstractListModel{parent},
    m_client{client}
    {}

    int rowCount(const QModelIndex& parent) const override
    {
        return m_lists.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        qDebug() << role;
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue(m_lists.at(index.row()).get());
    }

    void SetLists(const std::string& userID)
    {
        m_userID = userID;
        auto lists = m_client.GetLists(userID);
        for(auto& list : lists)
        {
            m_lists.emplace_back(std::make_unique<List>(list));
        }

        todolist::List list;
        list.Id = "lists/all";
        list.Title = "All";
        m_lists.emplace_back(std::make_unique<List>(list));
        selectList(QString::fromStdString(list.Id));
    }

    Q_INVOKABLE void addList(QString name)
    {
        qDebug() << "Add list" << name;
        m_client.AddList(m_userID, name.toStdString());
        todolist::List list;
        list.Id = "lists/" + name.toStdString();
        list.Title = name.toStdString();
        emit beginResetModel();
        m_lists.emplace_back(std::make_unique<List>(list));
        emit endResetModel();
        selectList(QString::fromStdString(list.Id));
    }

    Q_INVOKABLE void selectList(QString id)
    {
        qDebug() << "Select list" << id;
        m_currentList = id;
        emit listSelected(QString::fromStdString(m_userID), id);
    }

signals:
    void listSelected(QString, QString);

public slots:
    void addNewTask(todolist::Task task)
    {
        m_client.AddTask(m_currentList.toStdString(), task);
    }

private:
    TodoListService& m_client;
    std::vector<std::unique_ptr<List>> m_lists;
    QString m_currentList;
    std::string m_userID;
};

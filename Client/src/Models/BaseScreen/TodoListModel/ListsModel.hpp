#pragma once

#include <QObject>
#include <QAbstractListModel>

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
        auto lists = m_client.GetLists(userID);
        for(auto& list : lists)
        {
            m_lists.emplace_back(std::make_unique<List>(list));
        }

        todolist::List list;
        list.Id = "lists/all";
        list.Title = "All";
        m_lists.emplace_back(std::make_unique<List>(list));
    }

    Q_INVOKABLE void addList(QString name)
    {}

    Q_INVOKABLE void selectList(QString id)
    {
        qDebug() << "Select list" << id;
        emit listSelected(id);
    }

signals:
    void listSelected(QString);

private:
    TodoListService& m_client;
    std::vector<std::unique_ptr<List>> m_lists;
};

#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Chat.hpp"

#include <QDebug>

class ChatListModel : public QAbstractListModel
{
    std::vector<std::unique_ptr<QObject>> m_chats;
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
public:

    QString name() const { return "ChatListModel"; }

    ChatListModel(QObject* parent = nullptr)
        : QAbstractListModel{parent}
    {
        m_chats.emplace_back(new ChatInfo("First chat", parent));
        m_chats.emplace_back(new ChatInfo("Second chat", parent));
        m_chats.emplace_back(new ChatInfo("Third chat", parent));
    }

    int rowCount(const QModelIndex& parent) const override
    {
        return m_chats.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue( (m_chats.at(index.row()).get()) );
    }
};

#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Message/SimpleMessage.hpp"

#include <QDebug>

class ChatModel : public QAbstractListModel
{
    std::vector<std::unique_ptr<QObject>> m_messages;
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
public:

    QString name() const { return "ChatModel"; }

    ChatModel(QObject* parent = nullptr)
        : QAbstractListModel{parent}
    {
        m_messages.emplace_back(new SimpleMessage("First message", false, parent));
        m_messages.emplace_back(new SimpleMessage("Second message", true, parent));
        m_messages.emplace_back(new SimpleMessage("Third message", false, parent));
    }

    int rowCount(const QModelIndex& parent) const override
    {
        return m_messages.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue( (m_messages.at(index.row()).get()) );
    }
};

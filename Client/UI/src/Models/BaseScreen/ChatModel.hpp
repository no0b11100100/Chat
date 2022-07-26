#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Message/SimpleMessage.hpp"
#include "Header.hpp"

#include <QDebug>

class ChatModel : public QAbstractListModel
{
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
    Q_PROPERTY(QObject* header READ header CONSTANT)
    Q_PROPERTY(bool isChatSelected READ chatSelected NOTIFY chatSelectedChanged)

public:

    QString name() const { return "ChatModel"; }
    QObject* header() { return m_header.get(); }
    bool chatSelected() const { return m_currentChatID != ""; }

    ChatModel(QObject* parent = nullptr)
        : QAbstractListModel{parent},
        m_header{new Header(parent)}
    {}

    int rowCount(const QModelIndex& parent) const override
    {
        return m_messages.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        qDebug() << role;
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue( (m_messages.at(index.row()).get()) );
    }

public slots:
    void updateChatModel(const Header& header, QString currentChatID)
    {
        qDebug() << "updateChatModel for" << currentChatID;
        m_header->SetTitle(header.title());
        m_header->SetSecondLine(header.secondLine());
        updateMessageList(currentChatID);
    }

signals:
    void chatSelectedChanged();

private:
    std::vector<std::unique_ptr<QObject>> m_messages;
    std::unique_ptr<Header> m_header;
    QString m_currentChatID;

    void updateMessageList(QString currentChatID)
    {
        m_currentChatID = currentChatID;
        emit chatSelectedChanged();
        emit beginResetModel();
        // TODO: request message list for chat_id
        m_messages.clear();
        if (m_currentChatID == "1")
        {
            m_messages.emplace_back(new SimpleMessage("First message for first chat", false, this));
            m_messages.emplace_back(new SimpleMessage("Second message for first chat", true, this));
            m_messages.emplace_back(new SimpleMessage("Third message for first chat", false, this));
        }
        else if(m_currentChatID == "2")
        {
            m_messages.emplace_back(new SimpleMessage("First message for second chat", false, this));
            m_messages.emplace_back(new SimpleMessage("Second message for second chat", true, this));
            m_messages.emplace_back(new SimpleMessage("Third message for second chat", false, this));
        }
        else
        {
            m_messages.emplace_back(new SimpleMessage("First message for third chat", false, this));
            m_messages.emplace_back(new SimpleMessage("Second message for third chat", true, this));
            m_messages.emplace_back(new SimpleMessage("Third message for third chat", false, this));
        }
        emit endResetModel();
        // emit dataChanged(createIndex(0,0), createIndex(m_messages.size(), 0));
    }
};

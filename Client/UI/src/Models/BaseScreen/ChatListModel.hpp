#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Chat.hpp"
#include "Header.hpp"

#include <QDebug>

class ChatListModel : public QAbstractListModel
{
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
public:

    QString name() const { return "ChatListModel"; }

    ChatListModel(QObject* parent = nullptr)
        : QAbstractListModel{parent}
    {
        m_chats.emplace_back(new ChatInformation("1", "First chat", parent));
        m_chats.emplace_back(new ChatInformation("2", "Second chat", parent));
        m_chats.emplace_back(new ChatInformation("3", "Third chat", parent));
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

    Q_INVOKABLE void selectChat(QString chatID)
    {
        qDebug() << "Select chat by id" << chatID;
        auto chat = findChatByID(chatID);
        Header header;
        header.SetTitle(chat->title());
        // TODO: Add second line
        emit chatSelected(header, chatID);
    }

signals:
    void chatSelected(const Header&, QString);

public slots:
    void setLastMessage(QString chatID, QString message)
    {
        auto chat = findChatByID(chatID);
        chat->UpdateLastMessage(message);
    }

private:
    std::vector<std::shared_ptr<ChatInformation>> m_chats;

    std::shared_ptr<ChatInformation> findChatByID(QString id)
    {
        auto it = std::find_if(m_chats.begin(), m_chats.end(), [&](const std::shared_ptr<ChatInformation>& chat){
            return chat->id() == id;
        });

        if(it == m_chats.end()) qDebug() << "Cannot find chat for id" << id;

        return *it;
    }
};

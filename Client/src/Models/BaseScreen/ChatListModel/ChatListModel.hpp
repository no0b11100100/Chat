#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Chat.hpp"
#include "../ChatModel/Header.hpp"
#include "../../../../services/ChatService.hpp"

#include <QDebug>

class ChatListModel : public QAbstractListModel
{
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
public:

    QString name() const { return "ChatListModel"; }

    ChatListModel(ChatService& client, QObject* parent = nullptr)
        : QAbstractListModel{parent},
        m_client{client}
    {}

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
        emit chatSelected(header, chatID);
    }

    void SetChats(std::string userID)
    {
        auto chats = m_client.getUserChats(userID);
        for(const auto& chat : chats)
        {
            m_chats.emplace_back(std::make_shared<ChatInformation>(QString::fromStdString(chat.ChatID), QString::fromStdString(chat.Title), QString::fromStdString(chat.LastMessage)));
        }
    }

public slots:
    void updateLastMessage(QString chatID, QString message, QString time)
    {
        auto chat = findChatByID(chatID);
        chat->UpdateLastMessage(message, time);
    }

signals:
    void chatSelected(const Header&, QString);

private:
    std::vector<std::shared_ptr<ChatInformation>> m_chats;
    ChatService& m_client;

    std::shared_ptr<ChatInformation> findChatByID(QString id)
    {
        auto it = std::find_if(m_chats.begin(), m_chats.end(), [&](const std::shared_ptr<ChatInformation>& chat){
            return chat->id() == id;
        });

        if(it == m_chats.end()) qDebug() << "Cannot find chat for id" << id;

        return *it;
    }
};

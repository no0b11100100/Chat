#pragma once

#include "ChatModel.hpp"
#include "ChatListModel.hpp"
#include "../../../grpc_client/Client.hpp"

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel CONSTANT)
    Q_PROPERTY(QObject* chatListModel READ chatList CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(GRPCClient* client, QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(parent)},
          m_chatList{new ChatListModel(parent)},
          m_client{client}
    {
        QObject::connect(m_chatList.get(), &ChatListModel::chatSelected,this, &BaseScreen::setChat); //m_chatModel.get(), &ChatModel::updateChatModel);
    }

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }

    QString name() const { return "BaseScreen"; }

    void SetUser(user::Response userData)
    {
        auto chats = m_client->chatService().GetChats(userData.user_id());
        m_chatList->SetChats(chats);
    }


public slots:
    void setChat(const Header& header, QString chatID)
    {
        m_chatModel->SetHeader(header);
        auto messages = m_client->chatService().GetMessages(chatID.toStdString(), "", chat::Direction::Forward);
        m_chatModel->SetMessages(messages);
    }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
    GRPCClient* m_client;
};

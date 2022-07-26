#pragma once

#include "ChatModel.hpp"
#include "ChatListModel.hpp"

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel CONSTANT)
    Q_PROPERTY(QObject* chatListModel READ chatList CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(parent)},
          m_chatList{new ChatListModel(parent)}
    {
        QObject::connect(m_chatList.get(), &ChatListModel::chatSelected, m_chatModel.get(), &ChatModel::updateChatModel);
    }

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }

    QString name() const { return "BaseScreen"; }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
};

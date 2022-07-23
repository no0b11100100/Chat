#pragma once

#include "ChatModel.hpp"
#include "ChatListModel.hpp"

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel NOTIFY chatChanged)
    Q_PROPERTY(QObject* chatListModel READ chatList CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(parent)},
          m_chatList{new ChatListModel(parent)}
    {}

    QObject* chatModel() { return m_chatModel.get(); }
    QObject* chatList() { return m_chatList.get(); }

    QString name() const { return "BaseScreen"; }

signals:
    void chatChanged();

private:
    std::unique_ptr<ChatModel> m_chatModel;
    std::unique_ptr<ChatListModel> m_chatList;
};

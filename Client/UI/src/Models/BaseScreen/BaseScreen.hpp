#pragma once

#include "ChatModel.hpp"

class BaseScreen : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QObject* chatModel READ chatModel CONSTANT)
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    BaseScreen(QObject* parent = nullptr)
        : QObject{parent},
          m_chatModel{new ChatModel(parent)}
    {}

    QObject* chatModel() { return m_chatModel.get(); }

    QString name() const { return "BaseScreen"; }

private:
    std::unique_ptr<ChatModel> m_chatModel;
    // std::shared_ptr<QObject> m_chatList;
};

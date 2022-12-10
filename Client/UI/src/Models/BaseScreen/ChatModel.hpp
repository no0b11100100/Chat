#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Message/SimpleMessage.hpp"
#include "Header.hpp"
#include "../../../grpc_client/proto_gen/chat_service.pb.h"
// #include "../../../json/Types/Value/Value.h"

#include <QJsonDocument>
#include <QJsonObject>
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
        m_header{new Header(parent)},
        m_currentChatID{""}
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

    void SetHeader(const Header& header)
    {
        m_header->SetTitle(header.title());
    }

    void SaveSelectedChatID(QString chatID)
    {
        qDebug() << "Select chat" << chatID;
        m_currentChatID = chatID;
        emit chatSelectedChanged();
    }

    void SetMessages(const chat::Messages& messages)
    {
        emit beginResetModel();
        m_messages.clear();
        for(const auto& message : messages.messages())
        {
            //TODO: use own json
            auto s = message.message_json();
            QJsonDocument object = QJsonDocument::fromJson(QByteArray(s.data(), int(s.size())));
            QJsonObject message_json = object.object();
            std::string text = message_json["message"].toString().toStdString();
            m_messages.emplace_back(new SimpleMessage(QString::fromStdString(text), false));
        }

        emit endResetModel();
    }

    Q_INVOKABLE void sendMessage(QString message)
    {
        emit beginResetModel();
        m_messages.emplace_back(new SimpleMessage(message, true));
        emit endResetModel();
        emit sendingMessage(m_currentChatID, message);
    }

    void AddMessage(chat::ExchangedMessage message)
    {
        if(m_currentChatID != "" && QString::fromStdString(message.chat_id()) == m_currentChatID)
        {
            auto msg = message.message();
            auto s = msg.message_json();
            QJsonDocument object = QJsonDocument::fromJson(QByteArray(s.data(), int(s.size())));
            QJsonObject message_json = object.object();
            std::string text = message_json["message"].toString().toStdString();
            emit beginResetModel();
            m_messages.emplace_back(new SimpleMessage(QString::fromStdString(text), false));
            emit endResetModel();
        }
        else
        {
            qDebug() << "Not add message" << m_currentChatID << QString::fromStdString(message.chat_id());
        }
    }

signals:
    void chatSelectedChanged();
    void sendingMessage(QString, QString);

private:
    std::vector<std::unique_ptr<QObject>> m_messages;
    std::unique_ptr<Header> m_header;
    QString m_currentChatID;
};

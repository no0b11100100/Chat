#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <vector>
#include <memory>

#include "Message/SimpleMessage.hpp"
#include "Header.hpp"
#include "../../../services/ChatService.hpp"

#include "src/MultiMedia/multimedia.hpp"

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

    ChatModel(ChatService& client, QObject* parent = nullptr)
        : QAbstractListModel{parent},
        m_header{new Header(parent)},
        m_currentChatID{""},
        m_multimedia{new Multimedia()},
        m_client{client}
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
        m_header->SetCallAction([&]()
        {
            qDebug() << "Execute call action";
            auto status = m_client.CallTo(m_currentChatID.toStdString(), "");
            if (status == chat::CallStatus::Connected)
            {
                m_multimedia->audio()->SubscribeOnAudioInput([&](QByteArray data){
                    chat::CallData callData;
                    callData.Audio = QString(data).toStdString();
                    std::cout << "SubscribeOnAudioInput " << QString(data).toStdString() << std::endl;
                    m_client.SendCallData(callData);
                });

                m_client.handleNotifyCallData([&](chat::CallData data){
                    connect(this, &ChatModel::sendAudioStream, m_multimedia->audio(), &Audio::receiveStream);
                    QByteArray audio = QString::fromStdString(data.Audio).toUtf8();
                    emit sendAudioStream(audio);
                });

                m_multimedia->audio()->startStream();
            }
        }); //TODO
    }

    void SaveSelectedChatID(QString chatID)
    {
        qDebug() << "Select chat" << chatID;
        m_currentChatID = chatID;
        emit chatSelectedChanged();
    }

    void SetMessages(const std::vector<chat::Message>& messages)
    {
        emit beginResetModel();
        m_messages.clear();
        for(const auto& message : messages)
        {
            json s = message.MessageJSON;
            chat::TextMessage textMessage;
            textMessage = s;
            m_messages.emplace_back(new SimpleMessage(QString::fromStdString(textMessage.Text), false));
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

    void AddMessage(chat::Message message)
    {
        if(m_currentChatID != "" && QString::fromStdString(message.ChatID) == m_currentChatID)
        {
            json s = message.MessageJSON;
            chat::TextMessage textMessage;
            textMessage = s;
            emit beginResetModel();
            m_messages.emplace_back(new SimpleMessage(QString::fromStdString(textMessage.Text), false));
            emit endResetModel();
        }
        else
        {
            qDebug() << "Not add message" << m_currentChatID << QString::fromStdString(message.ChatID);
        }
    }

signals:
    void chatSelectedChanged();
    void sendingMessage(QString, QString);
    void sendAudioStream(QByteArray);

private:
    std::vector<std::unique_ptr<QObject>> m_messages;
    std::unique_ptr<Header> m_header;
    QString m_currentChatID;
    std::unique_ptr<Multimedia> m_multimedia;
    ChatService& m_client;
};

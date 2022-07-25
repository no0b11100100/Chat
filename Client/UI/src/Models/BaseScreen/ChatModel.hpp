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

public:

    QString name() const { return "ChatModel"; }
    QObject* header() { return m_header.get(); }

    ChatModel(QObject* parent = nullptr)
        : QAbstractListModel{parent},
        m_header{new Header(parent)}
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

public slots:
    void changeHeader(const Header& header)
    {
        m_header->SetTitle(header.title());
        m_header->SetSecondLine(header.secondLine());
    }

signals:
    void lastMessageChanhed(QString chatID, QString massage);

private:
    std::vector<std::unique_ptr<QObject>> m_messages;
    std::unique_ptr<Header> m_header;
};

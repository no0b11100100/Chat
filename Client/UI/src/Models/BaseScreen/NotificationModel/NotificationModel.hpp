#pragma once

#include <QObject>
#include <QAbstractListModel>
#include <QDebug>

#include "Notification.hpp"

class NotificationModel : public QAbstractListModel
{
    Q_OBJECT
    Q_PROPERTY(QString name READ name CONSTANT)
public:
    NotificationModel(QObject* parent = nullptr)
     : QAbstractListModel{parent}
    {
        // emit beginResetModel();
        m_notifications.emplace_back(new Notification("say hello", "User"));
        m_notifications.emplace_back(new Notification("say hello again", "User"));
        // emit endResetModel();
    }

    QString name() const { return "NotificationModel"; }

    int rowCount(const QModelIndex& parent) const override
    {
        return m_notifications.size();
    }

    QVariant data(const QModelIndex& index, int role=Qt::DisplayRole) const override
    {
        qDebug() << "NotificationModel data" << index.isValid() << (role != Qt::DisplayRole);
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();

        return QVariant::fromValue( (m_notifications.at(index.row()).get()) );
    }

    void AddNotification(const QString& message)
    {
        emit beginResetModel();
        m_notifications.emplace_back(new Notification(message, "User"));
        emit endResetModel();
    }

private:
    std::vector<std::unique_ptr<QObject>> m_notifications;
};

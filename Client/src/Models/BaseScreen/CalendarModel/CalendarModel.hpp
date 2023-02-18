#pragma once

#include <QObject>
#include <QDebug>

#include "../../../../services/CalendarService.hpp"

class CalendarModel : public QObject
{
    Q_OBJECT
public:
    CalendarModel(CalendarService& client, QObject* parent = nullptr)
    : QObject{parent},
    m_client{client}
    {
        QObject::connect(this, &CalendarModel::receiveMeeting, this, &CalendarModel::handleMeetingNotification);
        m_client.meetingUpdate([this](calendar::Meeting meeting){emit receiveMeeting(meeting);});
    }

    Q_INVOKABLE void createMeeting(QString name, QString participants)
    {
        qDebug() << "Create meeting" << name << "for" << participants;
        std::vector<std::string> v;
        v.push_back(participants.toStdString());
        m_client.CreateMeeting(name.toStdString(), v);
    }

public slots:
void handleMeetingNotification(calendar::Meeting meeting)
{
    qDebug() << "Recieve meeting" << QString::fromStdString(meeting.Title);
}

signals:
    void receiveMeeting(calendar::Meeting);

private:
    CalendarService& m_client;
};

#pragma once

#include <QObject>
#include <QDebug>

#include "../../../../services/CalendarService.hpp"
#include "Meeting.hpp"

class CalendarModel : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QList<Meeting*> meetings READ meetings NOTIFY meetingsChanged)
public:
    CalendarModel(CalendarService& client, QObject* parent = nullptr)
    : QObject{parent},
    m_client{client}
    {
        QObject::connect(this, &CalendarModel::receiveMeeting, this, &CalendarModel::handleMeetingNotification);
        m_client.meetingUpdate([this](calendar::Meeting meeting){emit receiveMeeting(meeting);});
    }

    Q_INVOKABLE void createMeeting(QString name, QString participants, QString startTime, QString endTime)
    {
        qDebug() << "Create meeting" << name << "for" << participants << startTime << "-" << endTime;
        std::vector<std::string> v;
        v.push_back(participants.toStdString());

        calendar::Meeting m;
        m.Title = name.toStdString();
        m.StartTime = startTime.toStdString();
        m.EndTime = endTime.toStdString();
        m.Participants = v;
        m_client.CreateMeeting(m);
        handleMeetingNotification(m);
    }

    QList<Meeting*> meetings() const {
        QList<Meeting*> result;
        for(auto& elementPointer : m_meetings)
            result.push_back(elementPointer.get());
        return result;
    }

public slots:
void handleMeetingNotification(calendar::Meeting meeting)
{
    qDebug() << "Recieve meeting" << QString::fromStdString(meeting.Title) << QString::fromStdString(meeting.StartTime) << QString::fromStdString(meeting.EndTime);
    m_meetings.emplace_back(std::make_unique<Meeting>(meeting));
    emit meetingsChanged();
}

signals:
    void receiveMeeting(calendar::Meeting);
    void meetingsChanged();

private:
    CalendarService& m_client;
    std::vector<std::unique_ptr<Meeting>> m_meetings;
};

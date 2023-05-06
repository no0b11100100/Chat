#pragma once

#include <cmath>

#include <QObject>
#include <QDebug>
#include <QDate>

#include "../../../../services/CalendarService.hpp"
#include "Meeting.hpp"
#include <QAbstractTableModel>

class CalendarModel : public QAbstractTableModel
{
    Q_OBJECT
    static constexpr int WEEK = 7;
    static constexpr int HOURS = 24;
    static constexpr int SECTIONS_PER_HOUR = 2;
public:
    CalendarModel(CalendarService& client, QObject* parent = nullptr)
    : QAbstractTableModel(parent),
    m_client{client}
    {
        QObject::connect(this, &CalendarModel::receiveMeeting, this, &CalendarModel::handleMeetingNotification);
        m_client.meetingUpdate([this](calendar::Meeting meeting){emit receiveMeeting(meeting);});

        m_meetings.reserve(7);
        for(int i = 0; i < 7; i++) m_meetings.emplace_back(std::vector<std::unique_ptr<Meeting>>{});
    }

    int rowCount(const QModelIndex & = QModelIndex()) const override { return HOURS * SECTIONS_PER_HOUR; }
    int columnCount(const QModelIndex & = QModelIndex()) const override { return WEEK; }

    QVariant data(const QModelIndex &index, int role) const override
    {
        if(!index.isValid() || role != Qt::DisplayRole)
            return QVariant();


        if (!m_meetings.at(index.column()).empty())
        {
            int h = index.row() / 2;
            int m = index.row() % 2 == 0 ? 0 : 30;
            for(const auto& meeting : m_meetings.at(index.column()))
            {
                auto timeParts = meeting->startTime().split(":");
                int meetingH = timeParts.at(0).toInt();
                int meetingM = timeParts.at(1).toInt();
                if(meetingH == h && m <= meetingM)
                {
                    qDebug() << index.row() << index.column() << h << m << meetingH << meetingM;
                    return QVariant::fromValue(meeting.get());
                }
            }
        }

        return QVariant();
    }

    QHash<int, QByteArray> roleNames() const override
    {
        return {{Qt::DisplayRole, "display"}};
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
        m.Date = QDate::currentDate().toString("dd.MM.yyyy").toStdString();
        // m_client.CreateMeeting(m);
        handleMeetingNotification(m);
    }

signals:
    void receiveMeeting(calendar::Meeting);

public slots:
void handleMeetingNotification(calendar::Meeting meeting)
{
    qDebug() << "Recieve meeting" << QString::fromStdString(meeting.Title) << QString::fromStdString(meeting.StartTime) << QString::fromStdString(meeting.EndTime);
    auto date = QDate::fromString(QString::fromStdString(meeting.Date), "dd.MM.yyyy");
    int dayCount = date.dayOfWeek();
    qDebug() << "Meeting day" << dayCount;
    qDebug() << m_meetings[dayCount-1].size();
    emit beginResetModel();
    m_meetings[dayCount-1].emplace_back(std::make_unique<Meeting>(meeting));
    emit endResetModel();
    qDebug() << "Add meeting";
}

private:
    CalendarService& m_client;
    std::vector<std::vector<std::unique_ptr<Meeting>>> m_meetings;

};

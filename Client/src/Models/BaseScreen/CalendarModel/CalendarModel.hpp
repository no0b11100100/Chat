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
                auto timeParts = meeting->endTime().split(":");
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

    QVariant headerData(int section, Qt::Orientation orientation, int role = Qt::DisplayRole) const override
    {
        qDebug() << section << orientation;
        if(orientation == Qt::Orientation::Horizontal)
        {
            QString data = getHeaderDataBySection(section);
            return data;
        }
        if(orientation == Qt::Orientation::Vertical)
        {
            return QString("%1:%2").arg(section / 2).arg(section % 2 == 0 ? 0 : 30);
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
        v.push_back(participants.toStdString()); // TODO: set userID(email)

        calendar::Meeting m;
        m.Title = name.toStdString();
        m.StartTime = startTime.toStdString();
        m.EndTime = endTime.toStdString();
        m.Participants = v;
        m.Date = QDate::currentDate().toString("dd.MM.yyyy").toStdString();
        m_client.CreateMeeting(m);
        handleMeetingNotification(m);
    }

    void SetCalendar(const std::string& userID)
    {
        auto currentDate = QDate::currentDate();
        QDate startWeek = currentDate;
        QDate endWeek = currentDate.addDays(7);
        if(currentDate.dayOfWeek() != 1)
        {
            startWeek = currentDate.addDays(-(currentDate.dayOfWeek()-1));
            endWeek = currentDate.addDays(7 - currentDate.dayOfWeek());
        }

        auto meetings = m_client.GetMeetings(userID, startWeek.toString("dd.MM.yyyy").toStdString(), endWeek.toString("dd.MM.yyyy").toStdString());
        for(const auto& meeting : meetings)
        {
            auto date = QDate::fromString(QString::fromStdString(meeting.Date), "dd.MM.yyyy");
            int dayCount = date.dayOfWeek();
            m_meetings[dayCount-1].emplace_back(std::make_unique<Meeting>(meeting));
        }
    }

signals:
    void receiveMeeting(calendar::Meeting);

public slots:
void handleMeetingNotification(calendar::Meeting meeting)
{
    qDebug() << "Recieve meeting" << QString::fromStdString(meeting.Title) << QString::fromStdString(meeting.StartTime) << QString::fromStdString(meeting.EndTime);
    auto date = QDate::fromString(QString::fromStdString(meeting.Date), "dd.MM.yyyy");

    if(date.weekNumber() != QDate::currentDate().weekNumber())
    {
        qDebug() << "Receive meeting not for current week, skip update";
        return;
    }

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

    QString getHeaderDataBySection(int section) const
    {
        auto currentDate = QDate::currentDate();
        int currentDay = currentDate.day();
        int positionInWeek = currentDate.dayOfWeek();

        qDebug() << currentDate;

        std::array<QString, 12> months = {"January","February","March","April","May","June","July","August","September","October","November","December"};
        std::array<QString, 7> days = {"Monday", "Tuesaday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"};

        QDate startWeek = currentDate;

        if(positionInWeek != 1) startWeek = currentDate.addDays(-(positionInWeek-1));

        startWeek = startWeek.addDays(section);

        return QString("%1 %2\n%3").arg(startWeek.day()).arg(months.at(startWeek.month()-1)).arg(days.at(section));

    }

};

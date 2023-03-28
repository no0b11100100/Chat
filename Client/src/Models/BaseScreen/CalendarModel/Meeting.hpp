#pragma once

#include <QObject>

#include "../../../../communication/calendar_client.hpp"

class Meeting : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title CONSTANT)
    Q_PROPERTY(QString startTime READ startTime CONSTANT)
    Q_PROPERTY(QString endTime READ endTime CONSTANT)
public:
    Meeting(calendar::Meeting meeting, QObject* parent = nullptr)
    : m_meeting{meeting}
    {}

    QString title() const { return QString::fromStdString(m_meeting.Title); }
    QString startTime() const { return QString::fromStdString(m_meeting.StartTime); }
    QString endTime() const { return QString::fromStdString(m_meeting.EndTime); }

private:
    calendar::Meeting m_meeting;
};
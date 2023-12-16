#pragma once

#include <string>
#include <QString>
#include <QTime>


static QString convertTime(std::string time)
{
    auto messageTime = QTime::fromString(QString::fromStdString(time), "hh:mm:ss.zzz");
    return QString("%1:%2").arg(messageTime.hour()).arg(messageTime.minute());
}

static QString monthNumberToMonthName(int number)
{
    std::array<QString, 12> months = {"January","February","March","April","May","June","July","August","September","October","November","December"};
    return months.at(number-1);
}

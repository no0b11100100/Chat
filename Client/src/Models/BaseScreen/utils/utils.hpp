#pragma once

#include <string>
#include <QString>
#include <QTime>


static QString convertTime(std::string time)
{
    auto messageTime = QTime::fromString(QString::fromStdString(time), "hh:mm:ss.zzz");
    return QString("%1:%2").arg(messageTime.hour()).arg(messageTime.minute());
}

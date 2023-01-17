#pragma once

#include <QTcpSocket>
#include <memory>
#include <iostream>
#include <functional>
#include <string>
#include <QString>
#include <QDebug>
#include <thread>
#include <QEventLoop>

namespace DefaultClient {
//https://stackoverflow.com/questions/2698145/how-do-i-execute-qtcpsocket-in-a-different-thread
class TCPClient : public QObject
{
    Q_OBJECT
    std::function<void(const std::string&)> m_reciverCallback;
    std::string m_addr;

public:
    TCPClient(std::string addr, std::function<void(const std::string&)> callback)
        : m_reciverCallback{callback},
        m_addr{addr}
    {}

    void connectToServer()
    {
        std::thread([this]
        {
            QEventLoop eventLoop;
            QTcpSocket* socket = new QTcpSocket(&eventLoop);

            QString addr = QString::fromStdString(m_addr);
            auto parts = addr.split(":");

            qDebug() << "Start server" << parts.at(0) << parts.at(1).toInt();
            socket->connectToHost(parts.at(0), parts.at(1).toInt());

            // enqueue or process the data
            QObject::connect(socket, &QTcpSocket::readyRead, &eventLoop, [this, socket]
            {
                qDebug() << "call readSocket";
                QByteArray data;
                data = socket->readAll();
                qDebug() << "Data from server" << data;
                QString str(data);
                handleMessage(str);
            });

            // Quit the loop (and thread) if the socket it disconnected. You could also try
            // reconnecting
            QObject::connect(socket, &QTcpSocket::disconnected, &eventLoop, [&eventLoop]
            {
                eventLoop.quit();
            });

            QObject::connect(this, &TCPClient::sendMessage, &eventLoop, [socket, &eventLoop](QString data)
            {
                QByteArray bytes = data.toUtf8();
                qDebug() << "before write" << bytes;
                socket->write(bytes);
                socket->waitForBytesWritten(1000);
                qDebug() << "afer write";
            });

            eventLoop.exec();

            delete socket;
        }).detach();
    }

    void sendPayload(const std::string& payload)
    {
        QString data = QString::fromStdString(payload) + "\n";
        emit sendMessage(data);
    }

    void handleMessage(QString data)
    {
        std::string response = data.toStdString();
        m_reciverCallback(response);
    }

signals:
    void sendMessage(QString);

};

};

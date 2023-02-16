#pragma once

#include <QTcpSocket>
#include <memory>
#include <iostream>
#include <functional>
#include <string>
#include <mutex>
#include <QString>
#include <QDebug>
#include <thread>
#include <QEventLoop>
#include <future>

namespace DefaultClient {
//https://stackoverflow.com/questions/2698145/how-do-i-execute-qtcpsocket-in-a-different-thread
class TCPClient : public QObject
{
    Q_OBJECT
    std::function<void(const std::string&)> m_reciverCallback;
    std::string m_addr;
    std::mutex m_mutex;
    QString m_incompletePackage;

public:
    TCPClient(std::string addr, std::function<void(const std::string&)> callback)
        : m_reciverCallback{callback},
        m_addr{addr}
    {}

    std::string init(std::string baseServiceAddr)
    {
        // std::promise<std::string> promise;
        std::string data = "";

        std::thread([&]
        {
            QTcpSocket* socket = new QTcpSocket();
            QString addr = QString::fromStdString(baseServiceAddr);
            auto parts = addr.split(":");
            qDebug() << "Start server" << parts.at(0) << parts.at(1).toInt();
            socket->connectToHost(parts.at(0), parts.at(1).toInt());

            qDebug() << "before write";
            socket->write("\n");
            socket->waitForBytesWritten(1000);
            qDebug() << "afer write";
            socket->waitForReadyRead(-1);
            QByteArray byteData = socket->readAll();
            socket->disconnectFromHost();
            QString strData(byteData);
            data = strData.toStdString();

            delete socket;
        }).join();

        // std::string data = promise.get_future().get();
        return data;
    }

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
                std::lock_guard<std::mutex> guard(m_mutex);
                qDebug() << "call readSocket" << socket->bytesAvailable();
                while(socket->bytesAvailable())
                {
                    QByteArray byteData = socket->readAll();
                    QString strData(byteData);
                    // qDebug() << "Data from server" << strData;
                    if (!m_incompletePackage.isEmpty())
                    {
                        strData = m_incompletePackage + strData;
                        m_incompletePackage = "";
                    }

                    auto packages = strData.split('\n');
                    qDebug() << "Packages size" << packages.size();
                    if(strData.back() != '\n')
                        m_incompletePackage = packages.takeLast();

                    for(auto message : packages)
                        if(!message.isEmpty())
                            handleMessage(message);
                }
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

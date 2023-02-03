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
                // QByteArray data;
                // data = socket->readAll();
                // QString str(data);
                // int index = str.indexOf("\n");
                // QString message = str.mid(0, index);
                // qDebug() << "Data from server" << message << "other" << str.mid(index+1, str.size());
                // // QString str(data);
                // handleMessage(message);
                // if (index != str.size() && str.mid(index+1, str.size()).back() == '\n')
                //     handleMessage(str.mid(index+1, str.size()-1));
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

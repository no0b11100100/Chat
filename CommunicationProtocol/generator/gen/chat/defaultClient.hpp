#pragma once

#include <QTcpSocket>
#include <memory>
#include <iostream>
#include <functional>
#include <string>
#include <QString>
#include <QDebug>


class Client : public QObject
{
    Q_OBJECT
    std::unique_ptr<QTcpSocket> m_connection;
    std::function<void(const std::string&)> m_reciverCallback;
public:
    Client(std::function<void(const std::string&)> callback)
    : m_connection{new QTcpSocket()},
    m_reciverCallback{callback}
    {
        QObject::connect(m_connection.get(), &QTcpSocket::readyRead, this, &Client::readSocket);
    }

    void connectToServer()
    {
        m_connection->connectToHost("127.0.0.1",8080);
        if(m_connection->waitForConnected(3000)) {
            std::cout << "Connected\n";
        } else {
            std::cout << "Can't connect\n";
            return;
        }
    }

    void disconnectFromServer()
    {
        if(m_connection->isOpen())
        {
            m_connection->close();
        }
    }

    void sendPayload(const std::string& payload)
    {
        QString data = QString::fromStdString(payload) + "\n";
        QByteArray bytes = data.toUtf8();
        qDebug() << "before write" << bytes;
        m_connection->write(bytes);
        m_connection->waitForBytesWritten(1000);
        qDebug() << "afer write";
    }

public slots:
    void readSocket()
    {
        qDebug() << "call readSocket";
        QByteArray data;
        data = m_connection->readAll();
        qDebug() << "Data from server" << data;
        QString str(data);
        std::string response = str.toStdString();
        m_reciverCallback(response);
    }
};

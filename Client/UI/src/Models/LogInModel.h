#pragma once
#include <functional>

#include <QObject>
#include <QDebug>


namespace Models {

class LogInModel : public QObject
{
    Q_OBJECT
public:
    LogInModel(std::function<void(QString, QString)> action, QObject* parent = nullptr)
        : QObject{parent},
          m_action{action}
    {}

    Q_INVOKABLE void triggerLogIn(QString email, QString password) {
        qDebug() << email << password;
        m_action(email, password);
    }

private:
    std::function<void(QString, QString)> m_action;

};

}

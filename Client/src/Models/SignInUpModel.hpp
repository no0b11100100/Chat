#pragma once

#include "../../services/UserService.hpp"

#include <functional>

#include <QObject>
#include <QDebug>

#include <string>

namespace Models {

using user::SignIn;
using user::SignUp;

class SignInUpModel : public QObject
{
    Q_OBJECT

    using SignInAction = std::function<std::string(SignIn)>;
    using SignUpAction = std::function<std::string(SignUp)>;

public:
    SignInUpModel(SignInAction signInAction, SignUpAction signUpAction, QObject* parent = nullptr)
        : QObject{parent},
          m_signInAction{signInAction},
          m_signUpAction{signUpAction}
    {}

    Q_INVOKABLE void signIn(QString email, QString password) {
        emit statusMessage("");
        qDebug() << email << password;
        SignIn data;
        data.Email = email.toStdString();
        data.Password = password.toStdString();
        std::string message = m_signInAction(data);
        emit statusMessage(message.c_str());
    }

    Q_INVOKABLE void signUp(QString name, QString nickName, QString email, QString password, QString confirmedPassword) {
        qDebug() << name << email << password << confirmedPassword;
        emit statusMessage("");
        SignUp data;
        data.Name = name.toStdString();
        data.NickName = nickName.toStdString();
        data.Email = email.toStdString();
        data.Password = password.toStdString();
        data.ConfirmedPassword = confirmedPassword.toStdString();

        std::string message = m_signUpAction(data);
        emit statusMessage(message.c_str());
    }

signals:
    void statusMessage(const QString& message);

private:
    SignInAction m_signInAction;
    SignUpAction m_signUpAction;
};

}

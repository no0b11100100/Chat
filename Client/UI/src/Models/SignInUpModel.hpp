#pragma once

#include "../../grpc_client/proto_gen/base_service.pb.h"

#include <functional>

#include <QObject>
#include <QDebug>

#include <string>

namespace Models {

class SignInUpModel : public QObject
{
    Q_OBJECT

    using SignInAction = std::function<std::string(base::SignIn)>;
    using SignUpAction = std::function<std::string(base::SignUp)>;

public:
    SignInUpModel(SignInAction signInAction, SignUpAction signUpAction, QObject* parent = nullptr)
        : QObject{parent},
          m_signInAction{signInAction},
          m_signUpAction{signUpAction}
    {}

    Q_INVOKABLE void signIn(QString email, QString password) {
        emit statusMessage("");
        qDebug() << email << password;
        base::SignIn data;
        data.set_email(email.toStdString());
        data.set_password(password.toStdString());
        std::string message = m_signInAction(data);
        emit statusMessage(message.c_str());
    }

    Q_INVOKABLE void signUp(QString name, QString nickName, QString email, QString password, QString confirmedPassword) {
        qDebug() << name << email << password << confirmedPassword;
        emit statusMessage("");
        base::SignUp data;
        data.set_name(name.toStdString());
        data.set_nickname(nickName.toStdString());
        data.set_email(email.toStdString());
        data.set_password(password.toStdString());
        data.set_confirmedpassword(confirmedPassword.toStdString());

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

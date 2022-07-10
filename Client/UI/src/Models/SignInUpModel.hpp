#pragma once

#include "../../grpc_client/proto_gen/chat.pb.h"

#include <functional>

#include <QObject>
#include <QDebug>

#include <string>

namespace Models {

class SignInUpModel : public QObject
{
    Q_OBJECT

    using SignInAction = std::function<std::string(chat::SignIn)>;
    using SignUpAction = std::function<void(chat::SignUp)>;

public:
    SignInUpModel(SignInAction signInAction, SignUpAction signUpAction, QObject* parent = nullptr)
        : QObject{parent},
          m_signInAction{signInAction},
          m_signUpAction{signUpAction}
    {}

    Q_INVOKABLE void signIn(QString email, QString password) {
        emit statusMessage("");
        qDebug() << email << password;
        chat::SignIn data;
        data.set_email(email.toStdString());
        data.set_password(password.toStdString());
        std::string message = m_signInAction(data);
        emit statusMessage(message.c_str());
    }

    Q_INVOKABLE void signUp(QString name, QString email, QString password, QString confirmedPassword) {
        qDebug() << name << email << password << confirmedPassword;
    }

signals:
    void statusMessage(const QString& message);

private:
    SignInAction m_signInAction;
    SignUpAction m_signUpAction;
};

}

#pragma once
#include <functional>

#include <QObject>
#include <QDebug>


namespace Models {

class SignInUpModel : public QObject
{
    Q_OBJECT

    using SignInAction = std::function<void(QString, QString)>;
    using SignUpAction = std::function<void(QString, QString, QString, QString)>;
public:
    SignInUpModel(SignInAction signInAction, SignUpAction signUpAction, QObject* parent = nullptr)
        : QObject{parent},
          m_signInAction{signInAction},
          m_signUpAction{signUpAction}
    {}

    Q_INVOKABLE void signIn(QString email, QString password) {
        qDebug() << email << password;
    }

    Q_INVOKABLE void signUp(QString name, QString email, QString password, QString confirmedPassword) {
        qDebug() << name << email << password << confirmedPassword;
    }

private:
    SignInAction m_signInAction;
    SignUpAction m_signUpAction;
};

}

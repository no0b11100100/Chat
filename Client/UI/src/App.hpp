#pragma once
#include <unordered_map>
#include <string>
#include <memory>

#include <QObject>

#include "../grpc_client/Client.hpp"
#include "Models/LogInModel.h"

constexpr const char* LOG_IN_MODEL = "logInModel";


class App : public QObject {
    Q_OBJECT

    Q_PROPERTY(QObject* model READ currentModel NOTIFY modelChanged)
public:
    App(QObject* parent = nullptr)
        : QObject{parent}
    {
        initModels(parent);
    }

    void initModels(QObject* parent) {
        auto signInAction = [&](QString email, QString password) -> void
        {
            Info i;
            i.email = email.toStdString();
            i.password = password.toStdString();
            m_grpcClient.baseService().LogIn(i);
        };

        auto signUpAction = [&](QString, QString, QString, QString) -> void
        {
            qDebug() << "signUp";
        };

        m_models[LOG_IN_MODEL] = std::make_unique<Models::SignInUpModel>(signInAction, signUpAction, parent);
    }

    QObject* currentModel()
    {
        return m_models[LOG_IN_MODEL].get();
    }

signals:
    void modelChanged();

private:
    GRPCClient m_grpcClient;

    std::unordered_map<std::string, std::unique_ptr<QObject>> m_models;
};

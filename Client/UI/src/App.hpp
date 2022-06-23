#pragma once
#include <unordered_map>
#include <string>
#include <memory>

#include <QObject>

#include "../grpc_client/Client.hpp"
#include "Models/SignInUpModel.hpp"

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
        auto signInAction = [&](SignIn data) -> void
        {
            m_grpcClient.baseService().signIn(data);
        };

        auto signUpAction = [&](SignUp data) -> void
        {
            m_grpcClient.baseService().signUp(data);
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

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
        : QObject{parent},
        m_userID{""}
    {
        initModels(parent);
    }

    void initModels(QObject* parent) {
        auto signInAction = [&](SignIn data) -> std::string
        {
            auto result = m_grpcClient.baseService().signIn(data);
            m_userID = result.user_id();
            return result.errormessage();
        };

        auto signUpAction = [&](SignUp data)
        {
            auto result = m_grpcClient.baseService().signUp(data);
            m_userID = result.user_id();
            return result.errormessage();
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
    std::string m_userID;
};

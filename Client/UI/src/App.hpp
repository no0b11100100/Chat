#pragma once
#include <unordered_map>
#include <string>
#include <memory>
#include <thread>
#include <chrono>

#include <QObject>

#include "../grpc_client/Client.hpp"
#include "Models/SignInUpModel.hpp"
#include "Models/BaseScreen/BaseScreen.hpp"

constexpr const char* LOG_IN_MODEL = "logInModel";
constexpr const char* BASE_SCREEN_MODEL = "baseScreenModel";

class App : public QObject {
    Q_OBJECT

    Q_PROPERTY(QObject* model READ currentModel NOTIFY modelChanged)
public:
    App(QObject* parent = nullptr)
        : QObject{parent},
        m_currentModel{nullptr},
        m_userID{""}
    {
        initModels(parent);
        m_currentModel = m_models[BASE_SCREEN_MODEL];
    }

    void initModels(QObject* parent) {
        m_models[LOG_IN_MODEL] = std::make_unique<Models::SignInUpModel>(
            [this](SignIn data){ return signInAction(data); },
            [this](SignUp data){ return signUpAction(data); },
            parent);
        m_models[BASE_SCREEN_MODEL] = std::make_unique<BaseScreen>(parent);
    }

    QObject* currentModel()
    {
        return m_currentModel.get();
    }

private:
    std::string signUpAction(SignUp data) {
        auto result = m_grpcClient.baseService().signUp(data);
        auto userID = result.user_id();
        if (userID != "")
        {
            m_userID = userID;
        }
        std::cout << "SignUp " << m_userID << std::endl;
        return result.errormessage();
    }

    std::string signInAction(SignIn data) {
        auto result = m_grpcClient.baseService().signIn(data);
        auto userID = result.user_id();
        if (userID != "")
        {
            m_userID = userID;
        }
        std::cout << "SignIn " << m_userID << std::endl;
        return result.errormessage();
    }

signals:
    void modelChanged();

private:
    GRPCClient m_grpcClient;
    std::unordered_map<std::string, std::shared_ptr<QObject>> m_models;
    std::shared_ptr<QObject> m_currentModel;
    std::string m_userID;
};

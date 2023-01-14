#pragma once
#include <unordered_map>
#include <string>
#include <memory>
#include <thread>
#include <chrono>

#include <QObject>

// #include "../grpc_client/Client.hpp"
#include "Models/SignInUpModel.hpp"
#include "Models/BaseScreen/BaseScreen.hpp"
#include "../services/Client.hpp"

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
        m_currentModel = m_models[LOG_IN_MODEL];
        QObject::connect(this, &App::activateBaseScreen, this, &App::changeToBaseScreen);
    }

    void initModels(QObject* parent) {
        m_models[LOG_IN_MODEL] = std::make_unique<Models::SignInUpModel>(
            [this](user::SignIn data){ return signInAction(data); },
            [this](user::SignUp data){ return signUpAction(data); },
            parent);
        m_models[BASE_SCREEN_MODEL] = std::make_unique<BaseScreen>(&m_Client, parent);
    }

    QObject* currentModel()
    {
        return m_currentModel.get();
    }

private:
    // TODO: investigate delay
    std::string signUpAction(user::SignUp data) {
        auto result = m_Client.userService().signUp(data);
        auto userID = result.UserID;
        if ((int)result.Status == 0) {
            m_userID = userID;
            std::thread([this, result](){
                std::this_thread::sleep_for(std::chrono::seconds(2));
                emit activateBaseScreen(result);
            }).detach();
        }
        return result.StatusMessage;
    }

    std::string signInAction(user::SignIn data) {
        auto result = m_Client.userService().signIn(data);
        auto userID = result.UserID;
        if ((int)result.Status == 0) {
            m_userID = userID;
            std::thread([this, result](){
                std::this_thread::sleep_for(std::chrono::seconds(2));
                emit activateBaseScreen(result);
            }).detach();
        }
        return result.StatusMessage;
    }

public slots:
    void changeToBaseScreen(user::Response userData)
    {
        m_currentModel = m_models[BASE_SCREEN_MODEL];
        BaseScreen* baseScreen = static_cast<BaseScreen*>(m_currentModel.get());
        baseScreen->SetUser(userData);
        emit modelChanged();
    }

signals:
    void modelChanged();
    void activateBaseScreen(user::Response);

private:
    Client m_Client;
    std::unordered_map<std::string, std::shared_ptr<QObject>> m_models;
    std::shared_ptr<QObject> m_currentModel;
    std::string m_userID;
};

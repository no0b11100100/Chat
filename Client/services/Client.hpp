#pragma once

#include "UserService.hpp"
#include "ChatService.hpp"
#include <iostream>

constexpr const char* USER_SERVICE_ADDRESS = "localhost:8080";
constexpr const char* CHAT_SERVICE_ADDRESS = "localhost:8080";

class Client final {
public:
    explicit Client()
        : m_userService{UserService(USER_SERVICE_ADDRESS)},
        m_chatService{ChatService(CHAT_SERVICE_ADDRESS)}
    {}

    UserService& userService() { return m_userService; }
    ChatService& chatService() { return m_chatService; }
private:
    UserService m_userService;
    ChatService m_chatService;
};

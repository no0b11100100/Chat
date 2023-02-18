#pragma once

#include "UserService.hpp"
#include "ChatService.hpp"
#include "CalendarService.hpp"
#include <iostream>

constexpr const char* USER_SERVICE_ADDRESS = "localhost:1234";
constexpr const char* CHAT_SERVICE_ADDRESS = "localhost:1235";
constexpr const char* CALENDAR_SERVICE_ADDRESS = "localhost:1236";

class Client final {
public:
    explicit Client()
        : m_userService{UserService(USER_SERVICE_ADDRESS)},
        m_chatService{ChatService(CHAT_SERVICE_ADDRESS)},
        m_calendarService{CalendarService(CALENDAR_SERVICE_ADDRESS)}
    {}

    UserService& userService() { return m_userService; }
    ChatService& chatService() { return m_chatService; }
    CalendarService& calendarService() { return m_calendarService; }
private:
    UserService m_userService;
    ChatService m_chatService;
    CalendarService m_calendarService;
};

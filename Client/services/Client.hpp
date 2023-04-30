#pragma once

#include "UserService.hpp"
#include "ChatService.hpp"
#include "CalendarService.hpp"
#include "TodoListService.hpp"
#include <iostream>

constexpr const char* USER_SERVICE_ADDRESS = "localhost:1234";
constexpr const char* CHAT_SERVICE_ADDRESS = "localhost:1235";
constexpr const char* CALENDAR_SERVICE_ADDRESS = "localhost:1236";
constexpr const char* TODO_LIST_SERVICE_ADDRESS = "localhost:1237";

class Client final {
public:
    explicit Client()
        : m_userService{UserService(USER_SERVICE_ADDRESS)},
        m_chatService{ChatService(CHAT_SERVICE_ADDRESS)},
        m_calendarService{CalendarService(CALENDAR_SERVICE_ADDRESS)},
        m_todoListService{TodoListService(TODO_LIST_SERVICE_ADDRESS)}
    {}

    UserService& userService() { return m_userService; }
    ChatService& chatService() { return m_chatService; }
    CalendarService& calendarService() { return m_calendarService; }
    TodoListService& todoListService() { return m_todoListService; }
private:
    UserService m_userService;
    ChatService m_chatService;
    CalendarService m_calendarService;
    TodoListService m_todoListService;

};

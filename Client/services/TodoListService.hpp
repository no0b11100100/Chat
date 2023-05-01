#pragma once

#include <memory>
#include "../communication/todolist_client.hpp"

class TodoListService final {
    std::unique_ptr<todolist::TodoListServiceStub> m_stub;
public:
    TodoListService(const std::string& addr)
    : m_stub{new todolist::TodoListServiceStub(addr)}
    {}

    ResponseStatus AddTask(std::string listID, todolist::Task task)
    {
        return m_stub->AddTask(listID, task);
    }

    std::vector<todolist::Task> GetTasks(std::string listID) {
        return m_stub->GetTasks(listID);
    }
};
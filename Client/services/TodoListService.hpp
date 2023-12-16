#pragma once

#include <memory>
#include "../communication/todolist_client.hpp"

class TodoListService final {
    std::unique_ptr<todolist::TodoListServiceStub> m_stub;
public:
    TodoListService(const std::string& addr)
    : m_stub{new todolist::TodoListServiceStub(addr)}
    {}

    todolist::TodoListReponse AddTask(std::string userID,std::string listID, todolist::Task task) {
        task.IsCompleted = false;
        return m_stub->AddTask(userID, listID, task);
    }

    std::vector<todolist::Task> GetTasks(std::string userID, std::string listID) {
        return m_stub->GetTasks(userID, listID);
    }

    std::vector<todolist::List> GetLists(std::string userID) {
        return m_stub->GetLists(userID);
    }

    todolist::TodoListReponse AddList(std::string userID, std::string listName) {
        todolist::List list;
        list.Title = listName;
        return m_stub->AddList(userID, list);
    }

    ResponseStatus SetTaskState(std::string userID, std::string listID, std::string taskID, bool state) {
        return m_stub->SetTaskState(userID, listID, taskID, state);
    }
};

#pragma once

#include "common.hpp"

#include "types.hpp"

namespace todolist {

// enums

// structs
struct Task : public Types::ClassParser {
  std::string Id;
  std::string Description;
  bool IsCompleted;
  virtual json toJson() const override {
    json js({});
    js["id"] = Id;
    js["description"] = Description;
    js["iscompleted"] = IsCompleted;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Id = std::string();
    else
      Id = static_cast<std::string>(js["id"]);
    if (js.isNull())
      Description = std::string();
    else
      Description = static_cast<std::string>(js["description"]);
    if (js.isNull())
      IsCompleted = bool();
    else
      IsCompleted = static_cast<bool>(js["iscompleted"]);
  }
};

struct List : public Types::ClassParser {
  std::string Id;
  std::string Title;
  std::vector<Task> Tasks;
  virtual json toJson() const override {
    json js({});
    js["id"] = Id;
    js["title"] = Title;
    js["tasks"] = Tasks;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Id = std::string();
    else
      Id = static_cast<std::string>(js["id"]);
    if (js.isNull())
      Title = std::string();
    else
      Title = static_cast<std::string>(js["title"]);
    if (js.isNull())
      Tasks = std::vector<Task>();
    else
      Tasks = static_cast<std::vector<Task>>(js["tasks"]);
  }
};

struct TodoListReponse : public Types::ClassParser {
  std::string Id;
  ResponseStatus Status;
  virtual json toJson() const override {
    json js({});
    js["id"] = Id;
    js["status"] = Status;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Id = std::string();
    else
      Id = static_cast<std::string>(js["id"]);
    if (js.isNull())
      Status = ResponseStatus();
    else
      Status = static_cast<ResponseStatus>(js["status"]);
  }
};

class TodoListServiceStub : public Common::Base {
public:
  TodoListServiceStub() = default;

  TodoListServiceStub(const std::string &addr) : Common::Base(addr) {}

  TodoListReponse AddTask(std::string userID, std::string listID, Task task) {
    Common::MessageData _message;
    _message.Payload = json::array({userID, listID, task});
    _message.Endpoint = "TodoListService.AddTask";
    return Request<TodoListReponse>(_message);
  }
  std::vector<Task> GetTasks(std::string userID, std::string listID) {
    Common::MessageData _message;
    _message.Payload = json::array({userID, listID});
    _message.Endpoint = "TodoListService.GetTasks";
    return Request<std::vector<Task>>(_message);
  }
  std::vector<List> GetLists(std::string userID) {
    Common::MessageData _message;
    _message.Payload = json::array({userID});
    _message.Endpoint = "TodoListService.GetLists";
    return Request<std::vector<List>>(_message);
  }
  TodoListReponse AddList(std::string userID, List newList) {
    Common::MessageData _message;
    _message.Payload = json::array({userID, newList});
    _message.Endpoint = "TodoListService.AddList";
    return Request<TodoListReponse>(_message);
  }
  ResponseStatus SetTaskState(std::string userID, std::string listID,
                              std::string taskID, bool state) {
    Common::MessageData _message;
    _message.Payload = json::array({userID, listID, taskID, state});
    _message.Endpoint = "TodoListService.SetTaskState";
    return Request<ResponseStatus>(_message);
  }

private:
};

}; // namespace todolist
#pragma once

#include "common.hpp"

#include "types.hpp"

namespace todolist {

// enums

// structs
struct Task : public Types::ClassParser {
  std::string Description;
  virtual json toJson() const override {
    json js({});
    js["description"] = Description;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Description = std::string();
    else
      Description = static_cast<std::string>(js["description"]);
  }
};

struct List : public Types::ClassParser {
  std::string Id;
  std::string Title;
  virtual json toJson() const override {
    json js({});
    js["id"] = Id;
    js["title"] = Title;
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
  }
};

class TodoListServiceStub : public Common::Base {
public:
  TodoListServiceStub() = default;

  TodoListServiceStub(const std::string &addr) : Common::Base(addr) {}

  ResponseStatus AddTask(std::string listID, Task task) {
    Common::MessageData _message;
    _message.Payload = json::array({listID, task});
    _message.Endpoint = "TodoListService.AddTask";
    return Request<ResponseStatus>(_message);
  }
  std::vector<Task> GetTasks(std::string listID) {
    Common::MessageData _message;
    _message.Payload = json::array({listID});
    _message.Endpoint = "TodoListService.GetTasks";
    return Request<std::vector<Task>>(_message);
  }
  std::vector<List> GetLists(std::string userID) {
    Common::MessageData _message;
    _message.Payload = json::array({userID});
    _message.Endpoint = "TodoListService.GetLists";
    return Request<std::vector<List>>(_message);
  }

private:
};

}; // namespace todolist
#pragma once

#include "common.hpp"

#include "types.hpp"

namespace user {

// enums

// structs
struct UserInfo : public Types::ClassParser {
  std::string UserID;
  std::string Name;
  std::string NickName;
  std::string Photo;
  std::vector<std::string> Chats;
  std::string Email;
  std::string Password;
  virtual json toJson() override {
    json js({});
    js["UserID"] = UserID;
    js["Name"] = Name;
    js["NickName"] = NickName;
    js["Photo"] = Photo;
    js["Chats"] = Chats;
    js["Email"] = Email;
    js["Password"] = Password;
    return js;
  }

  virtual void fromJson(json js) override {
    UserID = static_cast<std::string>(js["UserID"]);
    Name = static_cast<std::string>(js["Name"]);
    NickName = static_cast<std::string>(js["NickName"]);
    Photo = static_cast<std::string>(js["Photo"]);
    Chats = static_cast<std::vector<std::string>>(js["Chats"]);
    Email = static_cast<std::string>(js["Email"]);
    Password = static_cast<std::string>(js["Password"]);
  }
};

struct Response : public Types::ClassParser {
  UserInfo Info;
  ResponseStatus Status;
  std::string StatusMessage;
  virtual json toJson() override {
    json js({});
    js["Info"] = Info;
    js["Status"] = Status;
    js["StatusMessage"] = StatusMessage;
    return js;
  }

  virtual void fromJson(json js) override {
    Info = static_cast<UserInfo>(js["Info"]);
    Status = static_cast<ResponseStatus>(js["Status"]);
    StatusMessage = static_cast<std::string>(js["StatusMessage"]);
  }
};

struct SignIn : public Types::ClassParser {
  std::string Email;
  std::string Password;
  virtual json toJson() override {
    json js({});
    js["Email"] = Email;
    js["Password"] = Password;
    return js;
  }

  virtual void fromJson(json js) override {
    Email = static_cast<std::string>(js["Email"]);
    Password = static_cast<std::string>(js["Password"]);
  }
};

struct SignUp : public Types::ClassParser {
  std::string Name;
  std::string NickName;
  std::string Email;
  std::string Password;
  std::string ConfirmedPassword;
  std::string Photo;
  virtual json toJson() override {
    json js({});
    js["Name"] = Name;
    js["NickName"] = NickName;
    js["Email"] = Email;
    js["Password"] = Password;
    js["ConfirmedPassword"] = ConfirmedPassword;
    js["Photo"] = Photo;
    return js;
  }

  virtual void fromJson(json js) override {
    Name = static_cast<std::string>(js["Name"]);
    NickName = static_cast<std::string>(js["NickName"]);
    Email = static_cast<std::string>(js["Email"]);
    Password = static_cast<std::string>(js["Password"]);
    ConfirmedPassword = static_cast<std::string>(js["ConfirmedPassword"]);
    Photo = static_cast<std::string>(js["Photo"]);
  }
};

class UserServiceStub : public Common::Base {
public:
  UserServiceStub() = default;

  UserServiceStub(const std::string &addr) : Common::Base(addr) {}

  Response SignIn(SignIn data) {
    json args = json::array({data});
    Common::MessageData _message;
    _message.Endpoint = "UserService.SignIn";
    _message.Payload = args.dump();
    return Request(_message).GetData<Response>();
  }
  Response SignUp(SignUp data) {
    json args = json::array({data});
    Common::MessageData _message;
    _message.Endpoint = "UserService.SignUp";
    _message.Payload = args.dump();
    return Request(_message).GetData<Response>();
  }

private:
};

}; // namespace user
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
  virtual json toJson() const override {
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
    if (js.isNull())
      UserID = std::string();
    else
      UserID = static_cast<std::string>(js["UserID"]);
    if (js.isNull())
      Name = std::string();
    else
      Name = static_cast<std::string>(js["Name"]);
    if (js.isNull())
      NickName = std::string();
    else
      NickName = static_cast<std::string>(js["NickName"]);
    if (js.isNull())
      Photo = std::string();
    else
      Photo = static_cast<std::string>(js["Photo"]);
    if (js.isNull())
      Chats = std::vector<std::string>();
    else
      Chats = static_cast<std::vector<std::string>>(js["Chats"]);
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["Email"]);
    if (js.isNull())
      Password = std::string();
    else
      Password = static_cast<std::string>(js["Password"]);
  }
};

struct Response : public Types::ClassParser {
  UserInfo Info;
  ResponseStatus Status;
  std::string StatusMessage;
  virtual json toJson() const override {
    json js({});
    js["Info"] = Info;
    js["Status"] = Status;
    js["StatusMessage"] = StatusMessage;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Info = UserInfo();
    else
      Info = static_cast<UserInfo>(js["Info"]);
    if (js.isNull())
      Status = ResponseStatus();
    else
      Status = static_cast<ResponseStatus>(js["Status"]);
    if (js.isNull())
      StatusMessage = std::string();
    else
      StatusMessage = static_cast<std::string>(js["StatusMessage"]);
  }
};

struct SignIn : public Types::ClassParser {
  std::string Email;
  std::string Password;
  virtual json toJson() const override {
    json js({});
    js["Email"] = Email;
    js["Password"] = Password;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["Email"]);
    if (js.isNull())
      Password = std::string();
    else
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
  virtual json toJson() const override {
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
    if (js.isNull())
      Name = std::string();
    else
      Name = static_cast<std::string>(js["Name"]);
    if (js.isNull())
      NickName = std::string();
    else
      NickName = static_cast<std::string>(js["NickName"]);
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["Email"]);
    if (js.isNull())
      Password = std::string();
    else
      Password = static_cast<std::string>(js["Password"]);
    if (js.isNull())
      ConfirmedPassword = std::string();
    else
      ConfirmedPassword = static_cast<std::string>(js["ConfirmedPassword"]);
    if (js.isNull())
      Photo = std::string();
    else
      Photo = static_cast<std::string>(js["Photo"]);
  }
};

class UserServiceStub : public Common::Base {
public:
  UserServiceStub() = default;

  UserServiceStub(const std::string &addr) : Common::Base(addr) {}

  Response SignIn(SignIn data) {
    Common::MessageData _message;
    _message.Payload = json::array({data});
    _message.Endpoint = "UserService.SignIn";
    return Request<Response>(_message);
  }
  Response SignUp(SignUp data) {
    Common::MessageData _message;
    _message.Payload = json::array({data});
    _message.Endpoint = "UserService.SignUp";
    return Request<Response>(_message);
  }

private:
};

}; // namespace user
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
    js["userid"] = UserID;
    js["name"] = Name;
    js["nickname"] = NickName;
    js["photo"] = Photo;
    js["chats"] = Chats;
    js["email"] = Email;
    js["password"] = Password;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      UserID = std::string();
    else
      UserID = static_cast<std::string>(js["userid"]);
    if (js.isNull())
      Name = std::string();
    else
      Name = static_cast<std::string>(js["name"]);
    if (js.isNull())
      NickName = std::string();
    else
      NickName = static_cast<std::string>(js["nickname"]);
    if (js.isNull())
      Photo = std::string();
    else
      Photo = static_cast<std::string>(js["photo"]);
    if (js.isNull())
      Chats = std::vector<std::string>();
    else
      Chats = static_cast<std::vector<std::string>>(js["chats"]);
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["email"]);
    if (js.isNull())
      Password = std::string();
    else
      Password = static_cast<std::string>(js["password"]);
  }
};

struct Response : public Types::ClassParser {
  UserInfo Info;
  ResponseStatus Status;
  std::string StatusMessage;
  virtual json toJson() const override {
    json js({});
    js["info"] = Info;
    js["status"] = Status;
    js["statusmessage"] = StatusMessage;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Info = UserInfo();
    else
      Info = static_cast<UserInfo>(js["info"]);
    if (js.isNull())
      Status = ResponseStatus();
    else
      Status = static_cast<ResponseStatus>(js["status"]);
    if (js.isNull())
      StatusMessage = std::string();
    else
      StatusMessage = static_cast<std::string>(js["statusmessage"]);
  }
};

struct SignIn : public Types::ClassParser {
  std::string Email;
  std::string Password;
  virtual json toJson() const override {
    json js({});
    js["email"] = Email;
    js["password"] = Password;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["email"]);
    if (js.isNull())
      Password = std::string();
    else
      Password = static_cast<std::string>(js["password"]);
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
    js["name"] = Name;
    js["nickname"] = NickName;
    js["email"] = Email;
    js["password"] = Password;
    js["confirmedpassword"] = ConfirmedPassword;
    js["photo"] = Photo;
    return js;
  }

  virtual void fromJson(json js) override {
    if (js.isNull())
      Name = std::string();
    else
      Name = static_cast<std::string>(js["name"]);
    if (js.isNull())
      NickName = std::string();
    else
      NickName = static_cast<std::string>(js["nickname"]);
    if (js.isNull())
      Email = std::string();
    else
      Email = static_cast<std::string>(js["email"]);
    if (js.isNull())
      Password = std::string();
    else
      Password = static_cast<std::string>(js["password"]);
    if (js.isNull())
      ConfirmedPassword = std::string();
    else
      ConfirmedPassword = static_cast<std::string>(js["confirmedpassword"]);
    if (js.isNull())
      Photo = std::string();
    else
      Photo = static_cast<std::string>(js["photo"]);
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
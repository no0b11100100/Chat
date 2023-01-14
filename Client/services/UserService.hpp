#pragma once

#include <memory>
#include "../communication/user_client.hpp"

class UserService final {
    std::unique_ptr<user::UserServiceStub> m_stub;
public:
    UserService(const std::string& addr)
    : m_stub{new user::UserServiceStub(addr)}
    {}

    user::Response signIn(user::SignIn data) {
        return m_stub->SignIn(data);
    }

    user::Response signUp(user::SignUp data) {
        return m_stub->SignUp(data);
    }
};

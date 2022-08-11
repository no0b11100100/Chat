#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto_gen/user_service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using user::User;
using user::SignIn;
using user::SignUp;
using user::Response;


class UserService final {
public:
    UserService(std::shared_ptr<Channel> channel)
        : _stub(User::NewStub(channel)) {}

    Response signIn(const SignIn& request) {
        std::cout << "signIn call\n";
        Response reply;
        ClientContext context;

        Status status = _stub->signIn(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    Response signUp(const SignUp& request) {
        Response reply;
        ClientContext context;

        Status status = _stub->signUp(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

private:
    std::unique_ptr<User::Stub> _stub;
};

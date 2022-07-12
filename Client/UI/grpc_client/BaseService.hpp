#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto_gen/chat.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using chat::Base;
using chat::SignIn;
using chat::SignUp;
using chat::Result;


class BaseService final {
public:
    BaseService(std::shared_ptr<Channel> channel)
        : _stub(Base::NewStub(channel)) {}

    Result signIn(const SignIn& request) {
        std::cout << "signIn call\n";
        Result reply;
        ClientContext context;

        Status status = _stub->signIn(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    Result signUp(const SignUp& request) {
        Result reply;
        ClientContext context;

        Status status = _stub->signUp(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

private:
    std::unique_ptr<Base::Stub> _stub;
};

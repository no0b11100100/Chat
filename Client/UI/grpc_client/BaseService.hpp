#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto/gen/chat.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using chat::Base;
using chat::UserLogIn;
using chat::ID;

struct Info {
    std::string email;
    std::string password;
    std::string name;
};

class BaseService final {
    public:
    BaseService(std::shared_ptr<Channel> channel)
      : _stub(Base::NewStub(channel)) {}

    std::string LogIn(const Info& i) {
        std::cout << "LogIn call\n";
        UserLogIn request;
        request.set_email(i.email);
        request.set_name(i.name);
        request.set_password(i.password);

        ID reply;
        ClientContext context;

        Status status = _stub->LogIn(&context, request, &reply);

        if (status.ok()) {
            std::cout << reply.id() << std::endl;
            return reply.id();
        }

        std::cout << status.error_code() << ": " << status.error_message() << std::endl;
        return "Error";
    }

    std::string Register(Info i) {
        UserLogIn request;
        request.set_email(i.email);
        request.set_name(i.name);
        request.set_password(i.password);
        request.set_nickname("nickname");

        ID reply;
        ClientContext context;

        Status status = _stub->Register(&context, request, &reply);

        if (status.ok()) {
            std::cout << reply.id() << std::endl;
            return reply.id();
        }

        std::cout << status.error_code() << ": " << status.error_message() << std::endl;
        return "Error";
    }

    private:
        std::unique_ptr<Base::Stub> _stub;
};

class Test {
public:
    void test(){
        BaseService base(grpc::CreateChannel("localhost:8080", grpc::InsecureChannelCredentials()));
        Info logIn;
        logIn.email = "test@account.com";
        logIn.password = "12345";
        // logIn.name = "name";

        std::string reply = base.LogIn(logIn);
        std::cout << "Greeter received: " << reply << std::endl;
    }
};


//int main() {
//    BaseService base(grpc::CreateChannel("172.17.0.2:8080", grpc::InsecureChannelCredentials()));
//    Info logIn;
//    logIn.email = "email";
//    logIn.password = "password";
//    logIn.name = "name";

//    std::string reply = base.LogIn(logIn);
//    std::cout << "Greeter received: " << reply << std::endl;
//    return 0;
//}

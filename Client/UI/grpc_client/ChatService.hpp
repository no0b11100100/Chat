#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto_gen/chat_service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using chat::Chat;

class ChatService final
{
public:
    ChatService(std::shared_ptr<Channel> channel)
        : _stub(Chat::NewStub(channel))
    {}

private:
    std::unique_ptr<Chat::Stub> _stub;
};
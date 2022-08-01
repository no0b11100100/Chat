#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto_gen/chat_service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using chat::Chat;
using chat::UserID;
using chat::UserChats;
using chat::ChatID;
using chat::ChatInformation;
using chat::ParticipantInfo;
using chat::MessageChan;
using chat::Messages;
using chat::Direction;

// https://gist.github.com/ppLorins/d4484b61f12b2d87ac5c8d50d0808974
class ChatService final
{
public:
    ChatService(std::shared_ptr<Channel> channel)
        : _stub(Chat::NewStub(channel))
    {}

    UserChats GetUserChats(const std::string& userID)
    {
        std::cout << "getUserChats " << userID;
        ClientContext context;
        UserID request;
        request.set_user_id(userID);
        UserChats reply;

        Status status = _stub->getUserChats(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "getUserChats error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    ChatInformation GetChatInfo(const std::string& chatID)
    {
        std::cout << "getChatInfo " << chatID;
        ClientContext context;
        ChatID request;
        request.set_chat_id(chatID);
        ChatInformation reply;

        Status status = _stub->getChatInfo(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "getChatInfo error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    ParticipantInfo GetParticipantInfo(const std::string& userID)
    {
        std::cout << "getParticipantInfo " << userID;
        ClientContext context;
        UserID request;
        request.set_user_id(userID);
        ParticipantInfo reply;

        Status status = _stub->getParticipantInfo(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "getParticipantInfo error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    Messages GetMessages(const std::string& messageID, Direction direction)
    {
        std::cout << "getMessages " << messageID << " " << static_cast<int>(direction) << std::endl;
        ClientContext context;
        MessageChan request;
        request.set_message_id(messageID);
        request.set_direction(direction);
        Messages reply;

        Status status = _stub->getMessages(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "getMessages error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

private:
    std::unique_ptr<Chat::Stub> _stub;
};
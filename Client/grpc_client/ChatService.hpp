#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "proto_gen/chat_service.grpc.pb.h"

#include <QDebug>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using grpc::ClientReader;

using chat::ChatInfo;
using chat::Chats;
using chat::Chat;
using chat::Direction;
using chat::MessageChan;
using chat::MessageType;
using chat::Message;
using chat::Messages;
using chat::UserID;
using chat::MessageID;
using chat::ChatData;
using chat::ExchangedMessage;
using chat::ReadMessage;
using chat::ChatData;


using MessageNotificationCallback = std::function<void(ExchangedMessage)>;

// https://gist.github.com/ppLorins/d4484b61f12b2d87ac5c8d50d0808974
// https://github.com/grpc/grpc/blob/master/examples/cpp/route_guide/route_guide_client.cc
class ChatService final
{
public:
    ChatService(std::shared_ptr<Channel> channel)
        : _stub(Chat::NewStub(channel))
    {}

    Chats GetChats(std::string userID)
    {
        UserID request;
        Chats reply;
        ClientContext context;

        request.set_user_id(userID);

        Status status = _stub->getChats(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    Messages GetMessages(std::string chatID, std::string from, Direction direction)
    {
        MessageChan request;
        Messages reply;
        ClientContext context;

        request.set_message_id(from);
        request.set_direction(direction);
        request.set_chat_id(chatID);

        Status status = _stub->getMessages(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return reply;
    }

    void ReadingMessage(std::string chatID, std::string messageID)
    {
        ReadMessage request;
        chat::Status reply;
        ClientContext context;

        request.set_chat_id(chatID);
        request.set_message_id(messageID);

        Status status = _stub->readMessage(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

    }

    void EditChat(ChatData request)
    {
        chat::Status reply;
        ClientContext context;

        Status status = _stub->editChat(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        }

    }

    void SendMessage(std::string chatID, Message message)
    {
        ExchangedMessage request;
        chat::Status reply;
        ClientContext context;

        request.set_chat_id(chatID);
        Message* msg = request.mutable_message();
        *msg = message;

        Status status = _stub->sendMessage(&context, request, &reply);

        if (!status.ok()) {
            std::cout << "SignIn error " << status.error_code() << ": " << status.error_message() << std::endl;
        } else {
            qDebug() << "SendMessage response " << QString::fromStdString(reply.status());
        }
    }

    //Note: Run in separate thread
    //TODO: change argument(s)
    void MessagesUpdated(MessageNotificationCallback handler)
    {
        // std::thread t([this, handler](){
        google::protobuf::Empty request;
        ClientContext context;
        ExchangedMessage reply;

        std::unique_ptr< ClientReader<ExchangedMessage> > reader(_stub->recieveMessage(&context, request));
        while(reader->Read(&reply))
        {
            handler(reply);
        }

        Status status = reader->Finish();

        if (status.ok()) {
            std::cout << "recieveMessage rpc succeeded." << std::endl;
        } else {
            std::cout << "recieveMessage rpc failed." << status.error_code() << ": " << status.error_message() << std::endl;
        }
        // });
        // t.detach();
    }

    //Note: Run in separate thread
    //TODO: change argument(s)
    void ChatChanged()
    {
        google::protobuf::Empty request;
        ClientContext context;
        ChatData reply;

        std::unique_ptr< ClientReader<ChatData> > reader(_stub->chatChanged(&context, request));
        while(reader->Read(&reply))
        {}

        Status status = reader->Finish();

        if (status.ok()) {
            std::cout << "recieveMessage rpc succeeded." << std::endl;
        } else {
            std::cout << "recieveMessage rpc failed." << status.error_code() << ": " << status.error_message() << std::endl;
        }
    }

private:
    std::unique_ptr<Chat::Stub> _stub;
};
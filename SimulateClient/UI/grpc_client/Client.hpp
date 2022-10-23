#include "UserService.hpp"
#include "ChatService.hpp"
#include <iostream>

constexpr const char* BASE_SERVICE_ADDRESS = "localhost:8090";
constexpr const char* CHAT_SERVICE_ADDRESS = "localhost:8090";

class GRPCClient final {
public:
explicit GRPCClient()
: m_userService{UserService(grpc::CreateChannel(BASE_SERVICE_ADDRESS, grpc::InsecureChannelCredentials()))},
m_chatService{ChatService(grpc::CreateChannel(CHAT_SERVICE_ADDRESS, grpc::InsecureChannelCredentials()))}
{
// std::function<void(std::string)> v;
// v = [](std::string str){
//     std::cout << str << std::endl;
// };
// m_chatService.MessagesUpdated(v);
startServices();
}

UserService& baseService() { return m_userService; }
ChatService& chatService() { return m_chatService; }
private:
UserService m_userService;
ChatService m_chatService;

void startServices()
{
//TODO: run all notifications handling
}
};

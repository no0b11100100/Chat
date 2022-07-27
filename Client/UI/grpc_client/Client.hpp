#include "BaseService.hpp"
#include "ChatService.hpp"

constexpr const char* BASE_SERVICE_ADDRESS = "localhost:8080";
constexpr const char* CHAT_SERVICE_ADDRESS = "localhost:12345";

class GRPCClient final {
public:
    explicit GRPCClient()
        : m_baseService{BaseService(grpc::CreateChannel(BASE_SERVICE_ADDRESS, grpc::InsecureChannelCredentials()))},
          m_chatService{ChatService(grpc::CreateChannel(CHAT_SERVICE_ADDRESS, grpc::InsecureChannelCredentials()))}
    {}

    BaseService& baseService() { return m_baseService; }
    ChatService& chatService() { return m_chatService; }
private:
    BaseService m_baseService;
    ChatService m_chatService;
};

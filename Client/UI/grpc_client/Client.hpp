#include "BaseService.hpp"

constexpr const char* BASE_SERVICE_ADDRESS = "localhost:8080";

class GRPCClient final {
public:
    explicit GRPCClient()
        : m_baseService{BaseService(grpc::CreateChannel(BASE_SERVICE_ADDRESS, grpc::InsecureChannelCredentials()))}
    {}

    BaseService& baseService() { return m_baseService; }
private:
    BaseService m_baseService;
};

// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: chat.proto
#ifndef GRPC_chat_2eproto__INCLUDED
#define GRPC_chat_2eproto__INCLUDED

#include "chat.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_generic_service.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/client_context.h>
#include <grpcpp/impl/codegen/completion_queue.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace chat {

class Base final {
 public:
  static constexpr char const* service_full_name() {
    return "chat.Base";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::chat::ID* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>> AsyncLogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>>(AsyncLogInRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>> PrepareAsyncLogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>>(PrepareAsyncLogInRaw(context, request, cq));
    }
    virtual ::grpc::Status Register(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::chat::ID* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>> AsyncRegister(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>>(AsyncRegisterRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>> PrepareAsyncRegister(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>>(PrepareAsyncRegisterRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, std::function<void(::grpc::Status)>) = 0;
      virtual void LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void Register(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Register(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>* AsyncLogInRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>* PrepareAsyncLogInRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>* AsyncRegisterRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::chat::ID>* PrepareAsyncRegisterRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::chat::ID* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>> AsyncLogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>>(AsyncLogInRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>> PrepareAsyncLogIn(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>>(PrepareAsyncLogInRaw(context, request, cq));
    }
    ::grpc::Status Register(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::chat::ID* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>> AsyncRegister(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>>(AsyncRegisterRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>> PrepareAsyncRegister(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::chat::ID>>(PrepareAsyncRegisterRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, std::function<void(::grpc::Status)>) override;
      void LogIn(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, ::grpc::ClientUnaryReactor* reactor) override;
      void Register(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, std::function<void(::grpc::Status)>) override;
      void Register(::grpc::ClientContext* context, const ::chat::UserLogIn* request, ::chat::ID* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::chat::ID>* AsyncLogInRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::chat::ID>* PrepareAsyncLogInRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::chat::ID>* AsyncRegisterRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::chat::ID>* PrepareAsyncRegisterRaw(::grpc::ClientContext* context, const ::chat::UserLogIn& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_LogIn_;
    const ::grpc::internal::RpcMethod rpcmethod_Register_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status LogIn(::grpc::ServerContext* context, const ::chat::UserLogIn* request, ::chat::ID* response);
    virtual ::grpc::Status Register(::grpc::ServerContext* context, const ::chat::UserLogIn* request, ::chat::ID* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_LogIn() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLogIn(::grpc::ServerContext* context, ::chat::UserLogIn* request, ::grpc::ServerAsyncResponseWriter< ::chat::ID>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Register() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestRegister(::grpc::ServerContext* context, ::chat::UserLogIn* request, ::grpc::ServerAsyncResponseWriter< ::chat::ID>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_LogIn<WithAsyncMethod_Register<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_LogIn() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::chat::UserLogIn, ::chat::ID>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::chat::UserLogIn* request, ::chat::ID* response) { return this->LogIn(context, request, response); }));}
    void SetMessageAllocatorFor_LogIn(
        ::grpc::MessageAllocator< ::chat::UserLogIn, ::chat::ID>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::chat::UserLogIn, ::chat::ID>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* LogIn(
      ::grpc::CallbackServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Register() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::chat::UserLogIn, ::chat::ID>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::chat::UserLogIn* request, ::chat::ID* response) { return this->Register(context, request, response); }));}
    void SetMessageAllocatorFor_Register(
        ::grpc::MessageAllocator< ::chat::UserLogIn, ::chat::ID>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::chat::UserLogIn, ::chat::ID>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Register(
      ::grpc::CallbackServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_LogIn<WithCallbackMethod_Register<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_LogIn() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Register() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_LogIn() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLogIn(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Register() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestRegister(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_LogIn() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->LogIn(context, request, response); }));
    }
    ~WithRawCallbackMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* LogIn(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Register() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Register(context, request, response); }));
    }
    ~WithRawCallbackMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Register(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_LogIn : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_LogIn() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::chat::UserLogIn, ::chat::ID>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::chat::UserLogIn, ::chat::ID>* streamer) {
                       return this->StreamedLogIn(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_LogIn() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status LogIn(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedLogIn(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::chat::UserLogIn,::chat::ID>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Register : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Register() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::chat::UserLogIn, ::chat::ID>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::chat::UserLogIn, ::chat::ID>* streamer) {
                       return this->StreamedRegister(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Register() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Register(::grpc::ServerContext* /*context*/, const ::chat::UserLogIn* /*request*/, ::chat::ID* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedRegister(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::chat::UserLogIn,::chat::ID>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_LogIn<WithStreamedUnaryMethod_Register<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_LogIn<WithStreamedUnaryMethod_Register<Service > > StreamedService;
};

}  // namespace chat


#endif  // GRPC_chat_2eproto__INCLUDED

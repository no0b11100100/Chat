protoc --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import --proto_path=proto/ ./proto/*.proto
# /UI/grpc_client >>> ../../UI/third-party/vcpkg/packages/protobuf_x64-linux/tools/protobuf/protoc -I proto --cpp_out=proto/gen --grpc_out=proto/gen  --plugin=protoc-gen-grpc=../../UI/third-party/vcpkg/packages/grpc_x64-linux/tools/grpc/grpc_cpp_plugin proto/chat.proto

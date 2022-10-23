protoc --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import --proto_path=. ./*.proto
python3 -m grpc_tools.protoc -I . --python_out=api --grpc_python_out=api ./*.proto
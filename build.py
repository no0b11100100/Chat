import os
import shutil
import argparse


def remove_gen_folders_from_common():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    api_path = os.path.join(common_path, 'proto','api')
    gen_path = os.path.join(common_path, 'proto','gen')

    shutil.rmtree(api_path, ignore_errors=True)
    shutil.rmtree(gen_path, ignore_errors=True)


def remove_copied_folders_in_ui():
    current_path = os.path.dirname(os.path.abspath(__file__))
    ui_api_path = os.path.join(current_path, 'Client','UI','grpc_client','proto_gen')
    shutil.rmtree(ui_api_path, ignore_errors=True)


def remove_copied_folders_in_client():
    current_path = os.path.dirname(os.path.abspath(__file__))
    client_server_path = os.path.join(current_path, 'Client','Server')
    client_server_api_path = os.path.join(client_server_path, 'api')
    client_server_common_path = os.path.join(client_server_path, 'common')
    shutil.rmtree(client_server_api_path, ignore_errors=True)
    shutil.rmtree(client_server_common_path, ignore_errors=True)


def remove_copied_folders_in_server():
    current_path = os.path.dirname(os.path.abspath(__file__))
    server_common_path =  os.path.join(current_path, 'RemoteServer','common')
    server_structs_path = os.path.join(current_path, 'RemoteServer', 'structs')
    shutil.rmtree(server_common_path, ignore_errors=True)
    shutil.rmtree(server_structs_path, ignore_errors=True)



def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('--client', default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--ui', default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--server', default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--all', default=False, action=argparse.BooleanOptionalAction)
    return parser.parse_args()


def copy_client_gen_api():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    api_path = os.path.join(common_path, 'proto','api')

    #copy api for Golang
    client_server_path = os.path.join(current_path, 'Client','Server')
    client_server_api_path = os.path.join(client_server_path, 'api')

    if not os.path.exists(client_server_api_path):
        os.mkdir(client_server_api_path)

    shutil.copytree(api_path, client_server_api_path, dirs_exist_ok=True)


def generate_proto_server():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    proto_path = os.path.join(common_path, 'proto')

    client_server_grpc_gen_path = os.path.join(proto_path, 'api')
    if not os.path.exists(client_server_grpc_gen_path):
        os.mkdir(client_server_grpc_gen_path)

    os.chdir(proto_path)
    os.system('make proto-gen')


def copy_common(target_path:str):
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')

    path = os.path.join(current_path, target_path, 'common')

    if not os.path.exists(path):
        os.mkdir(path)

    print(f'Copy {common_path} to {path}')
    shutil.copytree(common_path, path, dirs_exist_ok=True)


def build_client():
    print("Build Client")
    generate_proto_server()
    copy_client_gen_api()
    copy_common(os.path.join('Client','Server'))
    current_path = os.path.dirname(os.path.abspath(__file__))
    local_server_path = os.path.join(current_path, 'Client','Server')
    os.chdir(local_server_path)
    os.system('docker build -t client-server .')
    remove_copied_folders_in_client()


def copy_server_gen_api():
    current_path = os.path.dirname(os.path.abspath(__file__))
    server_path = os.path.join(current_path, 'RemoteServer')
    structs_path = os.path.join(server_path, 'structs')
    if not os.path.exists(structs_path):
        os.mkdir(structs_path)

    common_path = os.path.join(current_path, 'common')
    api_path = os.path.join(common_path, 'proto','api')

    files = [f for f in os.listdir(api_path) if f.endswith('_service.pb.go')]

    for f in files:
        shutil.copy(os.path.join(api_path, f), structs_path)


def build_server():
    print("Build Server")
    generate_proto_server()
    copy_server_gen_api()
    copy_common(os.path.join('RemoteServer'))
    current_path = os.path.dirname(os.path.abspath(__file__))
    remote_server_path = os.path.join(current_path, 'RemoteServer')
    os.chdir(remote_server_path)
    os.system('docker compose build')
    remove_copied_folders_in_server()


def generate_ui_proto():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    proto_path = os.path.join(common_path, 'proto')
    client_ui_grpc_gen_path = os.path.join(proto_path, 'gen')
    if not os.path.exists(client_ui_grpc_gen_path):
        os.mkdir(client_ui_grpc_gen_path)

    ui_path = os.path.join(current_path, 'Client', 'UI')
    protoc_path = os.path.join(ui_path, 'third-party','vcpkg','packages','protobuf_x64-linux','tools','protobuf','protoc')
    plagin_path = os.path.join(ui_path, 'third-party','vcpkg','packages','grpc_x64-linux','tools','grpc','grpc_cpp_plugin')
    os.chdir(proto_path)
    os.system(f'{protoc_path} -I {proto_path} --cpp_out=gen --grpc_out=gen --plugin=protoc-gen-grpc={plagin_path} *.proto')


def copy_ui_gen_api():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    gen_path = os.path.join(common_path, 'proto','gen')
    ui_api_path = os.path.join(current_path, 'Client','UI','grpc_client','proto_gen')

    if not os.path.exists(ui_api_path):
        os.mkdir(ui_api_path)

    shutil.copytree(gen_path, ui_api_path, dirs_exist_ok=True)


def build_ui():
    print("Build UI")
    generate_ui_proto()
    copy_ui_gen_api()

    current_path = os.path.dirname(os.path.abspath(__file__))
    ui_path = os.path.join(current_path, 'Client','UI')
    ui_build_folder_path = os.path.join(ui_path, 'build')
    shutil.rmtree(ui_build_folder_path, ignore_errors=True)
    os.mkdir(ui_build_folder_path)
    os.chdir(ui_build_folder_path)

    buildsystem_path = os.path.join(ui_path, 'third-party','vcpkg','scripts','buildsystems','vcpkg.cmake')
    if os.system(f'cmake .. -DCMAKE_TOOLCHAIN_FILE={buildsystem_path} && make -j7') != 0:
        os._exit(1)
    remove_copied_folders_in_ui()


if __name__ == '__main__':
    # docker rm $(docker ps -a -q) && docker image prune

    args = parse_args()

    if args.all:
        build_client()
        build_ui()
        build_server()

    if args.client:
        build_client()
    if args.ui:
        build_ui()
    if args.server:
        build_server()

    remove_gen_folders_from_common()

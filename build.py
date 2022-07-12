import os
import shutil


def copy_common():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    client_path = os.path.join(current_path, 'Client')
    client_server_path = os.path.join(client_path, 'Server')
    client_server_common_path = os.path.join(client_server_path, 'common')
    server_path =  os.path.join(current_path, 'RemoteServer')
    server_common_path = os.path.join(server_path, 'common')

    if not os.path.exists(client_server_common_path):
        os.mkdir(client_server_common_path)

    print(f'Copy {common_path} to {client_server_common_path}')
    shutil.copytree(common_path, client_server_common_path, dirs_exist_ok=True)

    if not os.path.exists(server_common_path):
        os.mkdir(server_common_path)

    print(f'Copy {common_path} to {server_common_path}')
    shutil.copytree(common_path, server_common_path, dirs_exist_ok=True)


def generate_proto():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    proto_path = os.path.join(common_path, 'proto')

    #create folder to generate Golang proto
    client_server_grpc_gen_path = os.path.join(proto_path, 'api')
    if not os.path.exists(client_server_grpc_gen_path):
        os.mkdir(client_server_grpc_gen_path)

    os.chdir(proto_path)
    os.system('make proto-gen')

    #create folder to generate C++ proto
    client_ui_grpc_gen_path = os.path.join(proto_path, 'gen')
    if not os.path.exists(client_ui_grpc_gen_path):
        os.mkdir(client_ui_grpc_gen_path)

    ui_path = os.path.join(current_path, 'Client', 'UI')
    protoc_path = os.path.join(ui_path, 'third-party','vcpkg','packages','protobuf_x64-linux','tools','protobuf','protoc')
    plagin_path = os.path.join(ui_path, 'third-party','vcpkg','packages','grpc_x64-linux','tools','grpc','grpc_cpp_plugin')
    os.chdir(proto_path)
    os.system(f'{protoc_path} -I {proto_path} --cpp_out=gen --grpc_out=gen --plugin=protoc-gen-grpc={plagin_path} *.proto')


def copy_gen_api():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    api_path = os.path.join(common_path, 'proto','api')

    #copy api for Golang
    client_server_path = os.path.join(current_path, 'Client','Server')
    client_server_api_path = os.path.join(client_server_path, 'api')

    if not os.path.exists(client_server_api_path):
        os.mkdir(client_server_api_path)

    shutil.copytree(api_path, client_server_api_path, dirs_exist_ok=True)

    #copy api for C++
    gen_path = os.path.join(common_path, 'proto','gen')
    ui_api_path = os.path.join(current_path, 'Client','UI','grpc_client','proto_gen')

    if not os.path.exists(ui_api_path):
        os.mkdir(ui_api_path)

    shutil.copytree(gen_path, ui_api_path, dirs_exist_ok=True)


def remove_gen_folders_from_common():
    current_path = os.path.dirname(os.path.abspath(__file__))
    common_path = os.path.join(current_path, 'common')
    api_path = os.path.join(common_path, 'proto','api')
    gen_path = os.path.join(common_path, 'proto','gen')

    shutil.rmtree(api_path, ignore_errors=True)
    shutil.rmtree(gen_path, ignore_errors=True)


def build():
    current_path = os.path.dirname(os.path.abspath(__file__))

    #build UI
    ui_path = os.path.join(current_path, 'Client','UI')
    ui_build_folder_path = os.path.join(ui_path, 'build')
    shutil.rmtree(ui_build_folder_path, ignore_errors=True)
    os.mkdir(ui_build_folder_path)
    os.chdir(ui_build_folder_path)

    buildsystem_path = os.path.join(ui_path, 'third-party','vcpkg','scripts','buildsystems','vcpkg.cmake')
    os.system(f'cmake .. -DCMAKE_TOOLCHAIN_FILE={buildsystem_path} && make -j7')

    #build local server
    local_server_path = os.path.join(current_path, 'Client','Server')
    os.chdir(local_server_path)
    os.system('docker build -t client-server .')

    #build remote server
    remote_server_path = os.path.join(current_path, 'RemoteServer')
    os.chdir(remote_server_path)
    os.system('docker compose build')


def remove_copied_folders():
    current_path = os.path.dirname(os.path.abspath(__file__))
    client_server_path = os.path.join(current_path, 'Client','Server')
    client_server_api_path = os.path.join(client_server_path, 'api')
    client_server_common_path = os.path.join(client_server_path, 'common')
    ui_api_path = os.path.join(current_path, 'Client','UI','grpc_client','proto_gen')
    server_common_path =  os.path.join(current_path, 'RemoteServer','common')

    shutil.rmtree(client_server_api_path, ignore_errors=True)
    shutil.rmtree(client_server_common_path, ignore_errors=True)
    shutil.rmtree(ui_api_path, ignore_errors=True)
    shutil.rmtree(server_common_path, ignore_errors=True)


if __name__ == '__main__':
    generate_proto()
    copy_gen_api()
    remove_gen_folders_from_common()
    copy_common()
    build()
    remove_copied_folders()

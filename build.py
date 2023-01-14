import os
import shutil
import argparse


def remove_copied_folders_in_ui():
    current_path = os.path.dirname(os.path.abspath(__file__))
    ui_api_path = os.path.join(current_path, 'Client','grpc_client','proto_gen')
    shutil.rmtree(ui_api_path, ignore_errors=True)


def remove_copied_folders_in_server():
    current_path = os.path.dirname(os.path.abspath(__file__))
    server_api_path =  os.path.join(current_path, 'Server','api')
    shutil.rmtree(server_api_path, ignore_errors=True)


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('--client', default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--server', default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--all', default=False, action=argparse.BooleanOptionalAction)
    return parser.parse_args()


def generate_communication():
    path = os.path.dirname(os.path.abspath(__file__))
    path = os.path.join(path, 'CommunicationProtocol', 'generator')
    os.chdir(path)

    os.system("make qface-gen")


def copy_server_communication():
    gen_path = os.path.dirname(os.path.abspath(__file__))
    gen_path = os.path.join(gen_path, 'CommunicationProtocol', 'generator', 'gen')
    import glob

    print("\n\n")
    files = []
    for filename in glob.iglob(gen_path + '**/**', recursive=True):
        if filename.endswith('.go'):
            files.append(filename)

    server_api_path = os.path.dirname(os.path.abspath(__file__))
    server_api_path = os.path.join(server_api_path, 'Server', 'api')

    if not os.path.exists(server_api_path):
        os.mkdir(server_api_path)

    for file in files:
        shutil.copy(file, server_api_path)


def build_server():
    print("Build Server")
    generate_communication()
    # copy_server_communication()
    # current_path = os.path.dirname(os.path.abspath(__file__))
    # remote_server_path = os.path.join(current_path, 'Server')
    # os.chdir(remote_server_path)
    # os.system('docker compose build')
    # remove_copied_folders_in_server()


def build_ui():
    print("Build UI")

    current_path = os.path.dirname(os.path.abspath(__file__))
    ui_path = os.path.join(current_path, 'Client')
    ui_build_folder_path = os.path.join(ui_path, 'build')
    shutil.rmtree(ui_build_folder_path, ignore_errors=True)
    os.mkdir(ui_build_folder_path)
    os.chdir(ui_build_folder_path)

    if os.system(f'cmake .. && make -j7') != 0:
        os._exit(1)
    # remove_copied_folders_in_ui()


if __name__ == '__main__':
    # docker rm $(docker ps -a -q) && docker image prune

    args = parse_args()

    if args.all:
        build_ui()
        build_server()

    if args.client:
        build_ui()
    if args.server:
        build_server()

    # remove_gen_folders_from_common()

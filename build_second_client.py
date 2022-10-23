import shutil
import os


def copy_client():
    current_path = os.path.dirname(os.path.abspath(__file__))
    simulation_path = os.path.join(current_path, 'SimulateClient')
    client_path = os.path.join(current_path, 'Client')
    shutil.copytree(client_path, simulation_path, dirs_exist_ok=True)


def replacement(filepath, text, subs, flags=0):
    replacement = ""
    with open(filepath, 'r') as f:
        file = f.readlines()
        for line in file:
            line = line.strip()
            changes = line.replace(text, subs)
            replacement = replacement + changes + "\n"

        with open(filepath, 'w') as f:
            f.write(replacement)


def change_for_client():
    current_path = os.path.dirname(os.path.abspath(__file__))
    communicator_path = os.path.join(current_path, 'SimulateClient','Server','communicator','communicator.go')
    dockerfile_path = os.path.join(current_path, 'SimulateClient','Server','Dockerfile')
    cpp_client_path = os.path.join(current_path, 'SimulateClient','UI','grpc_client','Client.hpp')
    replacement(communicator_path, '":8080")', '":8090")')
    replacement(dockerfile_path, '/client-server', '/client-server_simulation')
    replacement(cpp_client_path, ':8080";', ':8090";')


def build():
    current_path = os.path.dirname(os.path.abspath(__file__))

    #build UI
    ui_path = os.path.join(current_path, 'SimulateClient','UI')
    ui_build_folder_path = os.path.join(ui_path, 'build')
    shutil.rmtree(ui_build_folder_path, ignore_errors=True)
    os.mkdir(ui_build_folder_path)
    os.chdir(ui_build_folder_path)

    buildsystem_path = os.path.join(ui_path, 'third-party','vcpkg','scripts','buildsystems','vcpkg.cmake')
    if os.system(f'cmake .. -DCMAKE_TOOLCHAIN_FILE={buildsystem_path} && make -j7') != 0:
        os._exit(1)

    #build local server
    local_server_path = os.path.join(current_path, 'SimulateClient','Server')
    os.chdir(local_server_path)
    os.system('docker build -t client-server_simulation .')


if __name__ == '__main__':
    copy_client()
    change_for_client()
    build()
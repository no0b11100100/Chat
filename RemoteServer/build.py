
import os
import shutil


def build():
    print("Copy common folder...")
    root = os.path.dirname(os.path.abspath(__file__))
    print("root path", root)
    root_common_path = os.path.join(root, 'common')
    common_path = root.removesuffix('RemoteServer')

    common_path = os.path.join(common_path, 'common')
    print(common_path)
    if not os.path.exists(common_path):
        os.mkdir(root_common_path)
    shutil.copytree(common_path, root_common_path, dirs_exist_ok=True)
    print('Copied')

    if os.path.exists(os.path.join(root_common_path, 'go.mod')):
        os.remove(os.path.join(root_common_path, 'go.mod'))
    listDir(root, True)
    change_goMod(os.path.join(root, 'go.mod'))

    os.system('docker build -t server .')

    shutil.rmtree(root_common_path, ignore_errors=True)
    listDir(root)
    change_goMod(os.path.join(root, 'go.mod'), True)


def listDir(path, isDockerBuild=False):
    for filename in os.listdir(path):
        f = os.path.join(path, filename)
        if os.path.isfile(f):
            if f.endswith('.go'):
                if isDockerBuild:
                    replace(f, old_string='\"common\"', new_string='\"Chat/RemoteServer/common\"')
                else:
                    replace(f, old_string='\"Chat/RemoteServer/common\"', new_string='\"common\"')
        elif os.path.isdir(f) and not f.endswith('common'):
            listDir(f, isDockerBuild)


def replace(file_path, old_string, new_string):
    file = open(file_path, "r")
    replacement = ""
    for line in file:
        origiLine = line
        line = line.strip()
        if old_string in line:
            changes = origiLine.replace(old_string, new_string)
            print("find: change to", changes)
            replacement = replacement + changes
            continue

        replacement = replacement + origiLine

    file.close()
    fout = open(file_path, "w")
    fout.write(replacement)
    fout.close()


def change_goMod(file_path, is_docker=False):
    file = open(file_path, "r")
    replacement = ""
    for line in file:
        origiLine = line
        line = line.strip()
        if 'replace common' in line or 'common v' in line:
            if is_docker:
                changes = origiLine.removeprefix('// ')
                replacement = replacement + changes
            else:
                changes = '// ' + origiLine
                print("find: change to", changes)
                replacement = replacement + changes
            continue

        replacement = replacement + origiLine

    file.close()
    fout = open(file_path, "w")
    fout.write(replacement)
    fout.close()


if __name__ == '__main__':
    build()

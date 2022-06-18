# /bin/bash

cd build

rm -rf * && \
cmake .. -DCMAKE_TOOLCHAIN_FILE=/home/drago/Desktop/Golang/Chat/Client/UI/third-party/vcpkg/scripts/buildsystems/vcpkg.cmake && \
make && \
./UI

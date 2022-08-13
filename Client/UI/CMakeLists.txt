cmake_minimum_required(VERSION 3.14)

project(UI VERSION 0.1 LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set (VCPKG_HOME "${CMAKE_SOURCE_DIR}/third-party/vcpkg")
set (CMAKE_TOOLCHAIN_FILE "${VCPKG_HOME}/scripts/buildsystems/vcpkg.cmake")

find_package(gRPC CONFIG REQUIRED)
find_package(Protobuf CONFIG REQUIRED)
find_package(absl CONFIG REQUIRED)
find_package(Threads REQUIRED)

find_package(QT NAMES Qt6 Qt5 COMPONENTS Core Quick REQUIRED)
find_package(Qt${QT_VERSION_MAJOR} COMPONENTS Core Quick REQUIRED)

aux_source_directory(${CMAKE_CURRENT_SOURCE_DIR}/grpc_client/proto_gen/ SUB_SOURCES)
aux_source_directory(${CMAKE_CURRENT_SOURCE_DIR}/json/ SUB_SOURCES)

set(PROJECT_SOURCES
        ${SUB_SOURCES}
        main.cpp
        qml/qml.qrc
        grpc_client/UserService.hpp
        grpc_client/ChatService.hpp
        grpc_client/Client.hpp

        json/Types/Base.hpp
        json/Types/Integer/Integer.h
        json/Types/Integer/Integer.cpp
        json/Types/Number/Number.h
        json/Types/Number/Number.cpp
        json/Types/Bool/Bool.h
        json/Types/Bool/Bool.cpp
        json/Types/String/String.h
        json/Types/String/String.cpp
        json/Types/Null/Null.h
        json/Types/Null/Null.cpp
        json/Types/Vector/Vector.h
        json/Types/Vector/Vector.cpp
        json/Types/Map/Map.h
        json/Types/Map/Map.cpp
        json/Types/Map/Pair.hpp
        json/Types/Iterator/Iteartor.h
        json/Types/Value/Dumper.hpp
        json/Types/Value/Separators.hpp
        json/Types/Value/Parser.hpp
        json/Types/Value/Value.h
        json/Types/Value/Value.cpp

        src/Models/SignInUpModel.hpp
        src/Models/BaseScreen/ChatModel.hpp
        src/Models/BaseScreen/BaseScreen.hpp
        src/Models/BaseScreen/ChatListModel.hpp
        src/Models/BaseScreen/Chat.hpp
        src/Models/BaseScreen/Header.hpp
        src/Models/BaseScreen/Message/SimpleMessage.hpp

        src/App.hpp
        Logger/Logger.hpp
)

if(${QT_VERSION_MAJOR} GREATER_EQUAL 6)
    qt_add_executable(UI
        MANUAL_FINALIZATION
        ${PROJECT_SOURCES}
    )
# Define target properties for Android with Qt 6 as:
#    set_property(TARGET UI APPEND PROPERTY QT_ANDROID_PACKAGE_SOURCE_DIR
#                 ${CMAKE_CURRENT_SOURCE_DIR}/android)
# For more information, see https://doc.qt.io/qt-6/qt-add-executable.html#target-creation
else()
    if(ANDROID)
        add_library(UI SHARED
            ${PROJECT_SOURCES}
        )
# Define properties for Android with Qt 5 after find_package() calls as:
#    set(ANDROID_PACKAGE_SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/android")
    else()
        add_executable(UI
          ${PROJECT_SOURCES}
        )
    endif()
endif()

target_compile_definitions(UI
  PRIVATE $<$<OR:$<CONFIG:Debug>,$<CONFIG:RelWithDebInfo>>:QT_QML_DEBUG>)
target_link_libraries(UI
  PRIVATE Qt${QT_VERSION_MAJOR}::Core Qt${QT_VERSION_MAJOR}::Quick absl::status gRPC::grpc++ gRPC::grpc++_unsecure gRPC::grpc++_reflection protobuf::libprotobuf)

set_target_properties(UI PROPERTIES
    MACOSX_BUNDLE_GUI_IDENTIFIER my.example.com
    MACOSX_BUNDLE_BUNDLE_VERSION ${PROJECT_VERSION}
    MACOSX_BUNDLE_SHORT_VERSION_STRING ${PROJECT_VERSION_MAJOR}.${PROJECT_VERSION_MINOR}
)

if(QT_VERSION_MAJOR EQUAL 6)
    qt_import_qml_plugins(UI)
    qt_finalize_executable(UI)
endif()
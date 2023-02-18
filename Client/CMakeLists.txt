cmake_minimum_required(VERSION 3.14)

project(UI VERSION 0.1 LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

#set(QT_MEDIA_BACKEND "ffmpeg")
set(QT_ASSUME_STDERR_HAS_CONSOLE 1)

#set (VCPKG_HOME "${CMAKE_SOURCE_DIR}/third-party/vcpkg")
#set (CMAKE_TOOLCHAIN_FILE "${VCPKG_HOME}/scripts/buildsystems/vcpkg.cmake")

#find_package(gRPC CONFIG REQUIRED)
#find_package(Protobuf CONFIG REQUIRED)
#find_package(absl CONFIG REQUIRED)
find_package(Threads REQUIRED)

find_package(QT NAMES Qt6 Qt5 COMPONENTS Core Quick REQUIRED)
find_package(Qt${QT_VERSION_MAJOR} COMPONENTS Core Quick Multimedia REQUIRED)


set(PROJECT_SOURCES
        main.cpp
        qml/qml.qrc

        services/UserService.hpp
        services/ChatService.hpp
        services/CalendarService.hpp
        services/Client.hpp

        src/Models/SignInUpModel.hpp
        src/Models/BaseScreen/ChatModel.hpp
        src/Models/BaseScreen/BaseScreen.hpp
        src/Models/BaseScreen/ChatListModel.hpp
        src/Models/BaseScreen/Chat.hpp
        src/Models/BaseScreen/Header.hpp
        src/Models/BaseScreen/Message/SimpleMessage.hpp

        src/Models/BaseScreen/NotificationModel/NotificationModel.hpp
        src/Models/BaseScreen/NotificationModel/Notification.hpp

        src/Models/BaseScreen/CalendarModel/CalendarModel.hpp

        src/App.hpp
        Logger/Logger.hpp


        communication/json/Bool/Bool.h
        communication/json/Bool/Bool.cpp
        communication/json/Integer/Integer.h
        communication/json/Integer/Integer.cpp
        communication/json/Iterator/Iteartor.h
        communication/json/Map/Map.h
        communication/json/Map/Map.cpp
        communication/json/Map/Pair.hpp
        communication/json/Null/Null.h
        communication/json/Null/Null.cpp
        communication/json/Number/Number.h
        communication/json/Number/Number.cpp
        communication/json/String/String.h
        communication/json/String/String.cpp
        communication/json/Value/Value.h
        communication/json/Value/Value.cpp
        communication/json/Value/Dumper.hpp
        communication/json/Value/Parser.hpp
        communication/json/Value/Separators.hpp
        communication/json/Vector/Vector.h
        communication/json/Vector/Vector.cpp
        communication/json/Base.hpp

        communication/chat_client.hpp
        communication/common.hpp
        communication/defaultClient.hpp
        communication/user_client.hpp
        communication/calendar_client.hpp

        src/MultiMedia/Audio/audio.h
        src/MultiMedia/multimedia.hpp

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
  PRIVATE Qt${QT_VERSION_MAJOR}::Core Qt${QT_VERSION_MAJOR}::Quick Qt${QT_VERSION_MAJOR}::Multimedia)

set_target_properties(UI PROPERTIES
    MACOSX_BUNDLE_GUI_IDENTIFIER my.example.com
    MACOSX_BUNDLE_BUNDLE_VERSION ${PROJECT_VERSION}
    MACOSX_BUNDLE_SHORT_VERSION_STRING ${PROJECT_VERSION_MAJOR}.${PROJECT_VERSION_MINOR}
)

if(QT_VERSION_MAJOR EQUAL 6)
    qt_import_qml_plugins(UI)
    qt_finalize_executable(UI)
endif()

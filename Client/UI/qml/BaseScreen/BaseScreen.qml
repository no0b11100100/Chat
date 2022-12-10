import QtQuick 2.0

import "ChatListModel"
import "ChatModel"

Rectangle {
    id: root
    property var model

    Row {
        spacing: 1
        SideBar {
            id: sidebar
            width: 30
            height: root.height
            buttons: [{"text":"chat"}, {"text":"mail"}]
        }
        ChatListModel {
            id: chats
            height: root.height
            width: root.width / 5 - sidebar.width
            model: root.model.chatListModel
            Component.onCompleted: {
                console.log("Create ChatListModel", root.model.chatListModel === undefined, model.name, root.model.name)
            }
        }

        ChatModel {
            height: root.height
            width: root.width - chats.width - sidebar.width
            model: root.model.chatModel
            Component.onCompleted: {
                console.log("Create ChatModel", root.model.chatModel === undefined, model.name, root.model.name)
            }
        }
    }
}

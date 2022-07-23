import QtQuick 2.0

Rectangle {
    id: root
    property var model

    Row {
        spacing: 1
        ChatListModel {
            id: chats
            height: root.height
            width: root.width / 5
            model: root.model.chatListModel
            Component.onCompleted: {
                console.log("Create ChatListModel", root.model.chatListModel === undefined, model.name, root.model.name)
            }
        }

        ChatModel {
            height: root.height
            width: root.width - chats.width
            model: root.model.chatModel
            Component.onCompleted: {
                console.log("Create ChatModel", root.model.chatModel === undefined, model.name, root.model.name)
            }
        }
    }
}
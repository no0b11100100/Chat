import QtQuick 2.0

Rectangle {
    id: root
    property var model

    ChatModel {
        height: root.height
        width: root.width
        model: root.model.chatModel
        Component.onCompleted: {
            console.log("Create ChatModel", root.model.chatModel === undefined, model.name, root.model.name)
        }
    }
}
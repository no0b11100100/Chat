import QtQuick 2.0

Rectangle {
    id: root
    property var model

    width: item.width
    height: item.height

    Rectangle {
        id: item
        width: root.width
        height: background.height

        Rectangle {
            id: background
            anchors.right: root.model.sendByMe ? parent.right : undefined
            radius: 5
            color: "lightblue"
            height: text.height + 5
            width: text.width + 10
            Text {
                id: text
                text: root.model.message
                anchors.centerIn: parent
            }
        }
    }

    Component.onCompleted: {
        console.log(root.model === undefined, root.model.name)
    }
}
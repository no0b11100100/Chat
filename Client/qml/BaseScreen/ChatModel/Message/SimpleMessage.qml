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
            anchors.rightMargin: root.model.sendByMe ? 5 : 0
            radius: 5
            color: "lightblue"
            height: text.height + 5 + time.height
            width: text.width + 10 + time.width
            Text {
                id: text
                text: root.model.message // + "\n" + root.model.time
                anchors.left: background.left
            }
            Text {
                id: time
                text: root.model.time
                anchors.right: text.right
                anchors.bottom: background.bottom
            }
        }
    }

    Component.onCompleted: {
        console.log(root.model === undefined, root.model.name)
    }
}

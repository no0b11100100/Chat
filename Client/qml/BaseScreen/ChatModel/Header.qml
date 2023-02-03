import QtQuick 2.0
import QtQuick.Controls 2.0

Rectangle {
    id: root
    property var model
    // width: root.width
    height: chatInfo.height

    Column {
        anchors.left: root.left
        id: chatInfo
        Text {
            text: root.model === undefined ? "" : root.model.title
        }

        Text {
            text: root.model === undefined ? "" : root.model.secondLine
        }
    }
    Button {
        anchors.right: root.right
        text: "call"
        height: root.height
        onClicked: {
            root.model.call()
        }
    }
}

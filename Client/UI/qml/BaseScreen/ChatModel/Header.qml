import QtQuick 2.0

Rectangle {
    id: root
    property var model
    // width: root.width
    height: chatInfo.height

    Column {
        id: chatInfo
        Text {
            text: root.model.title
        }

        Text {
            text: root.model.secondLine
        }
    }
}
import QtQuick 2.0

Rectangle {
    id: root
    property var model

    Rectangle {
        anchors.fill: parent
        radius: 5
        border.color: "black"
        border.width: 1

        Text {
            text: root.model.title
            // anchors.centerIn: parent
        }
    }
}
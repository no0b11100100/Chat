import QtQuick 2.15

Rectangle {
    id: root
    property var model

    ListView {
        anchors.fill: parent
        model: root.model
        delegate: Rectangle {
            Text {
                text: modelData.name
            }
        }
    }
}
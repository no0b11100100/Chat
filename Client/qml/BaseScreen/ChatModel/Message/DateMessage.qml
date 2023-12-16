import QtQuick 2.0

Rectangle {
    id: root
    property var model
    Text {
        anchors.centerIn: parent
        text: root.model.message
    }
}

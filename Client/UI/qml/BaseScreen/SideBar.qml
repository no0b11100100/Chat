import QtQuick 2.0

Rectangle {
    property var buttons: []
    id: root

    Column {
        anchors.fill:parent
        Repeater {
            model: root.buttons
            Rectangle {
                width: root.width
                height: 30
                Text {
                    text: modelData.text
                }
                color:"green"
            }
        }
    }
}

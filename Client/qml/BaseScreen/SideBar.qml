import QtQuick 2.0

Rectangle {
    property var buttons: []
    property var action
    id: root

    Column {
        anchors.fill:parent
        spacing: 5
        Repeater {
            model: root.buttons
            Rectangle {
                width: root.width
                height: 30
                Text {
                    text: modelData.text
                }
                color:"green"

                MouseArea {
                    anchors.fill: parent
                    onClicked: {
                        root.action(modelData.id)
                    }
                }
            }
        }
    }
}

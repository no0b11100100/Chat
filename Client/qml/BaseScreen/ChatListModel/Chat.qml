import QtQuick 2.0

Rectangle {
    id: root
    property var model
    property var selectChat

    Rectangle {
        anchors.fill: parent
        radius: 5
        border.color: "black"
        border.width: 1

        Column {
            Row {
                spacing: 40
                Text {
                    text: root.model.title
                    // anchors.centerIn: parent
                }
                Text {
                    text: root.model.lastMessageTime
                }
            }
            Text {
                text: root.model.lastMessage
            }
        }

        MouseArea{
            anchors.fill: parent

            onClicked: {
                selectChat(model.id)
            }
        }
    }
}

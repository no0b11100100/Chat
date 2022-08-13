import QtQuick 2.0
import QtQuick.Controls

Rectangle {
    id: root
    property var action

    height: input.height

    Row {

        TextField {
            id: input
            placeholderText: "Enter the message"
            placeholderTextColor: "grey"
            width: root.width - trigger.width
            visible: root.visible
            color: "black"
            font.pointSize: 16

            Component.onCompleted: {
                input.background.color = "white"
                input.background.opacity = 0.7
            }
        }

        Button {
            id: trigger
            height: root.height
            text: "Send"
            visible: root.visible

            onClicked: {
                root.action(input.text)
            }
        }

    }
}
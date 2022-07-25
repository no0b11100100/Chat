import QtQuick 2.0

import "Message"

Rectangle {
    id: root
    property var model

    Column {
        spacing: 2
        Rectangle {
            id: header
            width: root.width
            height: chatInfo.height

            Column {
                id: chatInfo
                Text {
                    text: root.model.header.title
                }

                Text {
                    text: root.model.header.secondLine
                }
            }
        }

        ListView {
            width: root.width
            height: root.height - header.height
            model: root.model
            spacing: 2
            delegate: SimpleMessage{
                width: root.width
                model: display
            }

            visible: true // TODO

            Component.onCompleted: {
                console.log(root.model === undefined, root.model.name)
            }
        }

        Rectangle {
            visible: false //TODO
            width: root.width

            Text {
                text: "Please select chat"
                anchors.centerIn: parent
            }
        }
    }
}
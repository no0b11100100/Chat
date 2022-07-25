import QtQuick 2.0

import "Message"

Rectangle {
    id: root
    property var model

    Column {
        spacing: 2

        Header {
            id: header
            width: root.width
            model: root.model.header
        }

        ListView {
            width: root.width
            height: root.height - header.height - userInput.height
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

        UserInput {
            id: userInput
            width: root.width
        }
    }
}
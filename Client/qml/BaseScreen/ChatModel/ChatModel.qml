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
            model: root.model == undefined ? undefined : root.model.header
        }

        ListView {
            id: listView
            width: root.width
            height: root.height - header.height - userInput.height
            model: root.model
            spacing: 2
            delegate: SimpleMessage{
                width: root.width
                model: display
            }

            visible: root.model == undefined ? true : root.model.isChatSelected

            Component.onCompleted: {
                console.log(root.model === undefined, root.model === undefined ? "UNDEFINED" : root.model.name)
            }
        }

        Rectangle {
            visible: root.model == undefined ? true : !root.model.isChatSelected
            width: root.width
            height: root.height

            Text {
                text: "Please select chat"
                anchors.centerIn: parent
            }
        }

        UserInput {
            id: userInput
            width: root.width
            visible: listView.visible
            action: root.model.sendMessage
        }
    }
}

import QtQuick 2.0

Rectangle {
    id: root
    property var model

    Column{
        Row {
            Text {
                id: senderName
                text: root.model.senderName
            }

            Text {
                id: message
                text: root.model.text
            }
        }

        // Image {
        //     id: senderAvatar
        //     source: root.model.senderAvatar
        // }

        // Text{
        //     id: time
        //     text: root.model.time
        // }
    }

}

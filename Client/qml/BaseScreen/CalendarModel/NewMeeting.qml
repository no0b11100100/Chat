import QtQuick 2.15
import QtQuick.Window 2.15
import QtQuick.Controls 2.15

Window {
    id: root
    width: 600
    height: 600
    visible: true
    property var createMeeting

    Column {
        TextField {
            id: name
            placeholderText: "Meeting name"
        }
        TextField {
            id: participants
            placeholderText: "Participants"
        }
        Button {
            text: "Create"
            onClicked: {
                root.createMeeting(name.text, participants.text)
            }
        }
    }

}

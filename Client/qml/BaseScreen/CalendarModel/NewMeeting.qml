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
                root.createMeeting(name.text, participants.text, startTime.currentText, endTime.currentText)
            }
        }
        // Row {
        ComboBox {
            id: startTime
            editable: true
            Component.onCompleted: {
                var data = []
                const hours = ["1","2","3","4","5","6","7","8","9"]
                const minutes = ["00", "15", "30", "45"]
                for (var i = 0; i < hours.length; i += 1) {
                    for(var j = 0; j < minutes.length; j += 1) {
                        data.push(hours[i] + ":" + minutes[j])
                    }
                }
                console.log("Time data", data)
                startTime.model = data
            }
        }

        ComboBox {
            id: endTime
            Component.onCompleted: {
                var data = []
                const hours = ["1","2","3","4","5","6","7","8","9"]
                const minutes = ["00", "15", "30", "45"]
                for (var i = 0; i < hours.length; i += 1) {
                    for(var j = 0; j < minutes.length; j += 1) {
                        data.push(hours[i] + ":" + minutes[j])
                    }
                }
                console.log("Time data", data)
                endTime.model = data
            }
        }
    }

}

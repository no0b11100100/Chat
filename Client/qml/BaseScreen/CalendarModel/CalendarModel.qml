import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root
    property var model

    Column {
        Button{
            text: "New Meeting"
            onClicked: {
                console.log("Press create new meeting")
                var component = Qt.createComponent("qrc:/qml/BaseScreen/CalendarModel/NewMeeting.qml")
                var window    = component.createObject()
                window.createMeeting = root.model.createMeeting
                window.show()
            }
        }
        Date {
            id: date
            width: root.width
            height: 50
            model: [
                {"date":"1 december", "day":"Monday"},
                {"date":"2 december", "day":"Tuesaday"},
                {"date":"3 december", "day":"Wednesday"},
                {"date":"4 december", "day":"Thursday"},
                {"date":"5 december", "day":"Friday"},
                {"date":"6 december", "day":"Saturday"},
                {"date":"7 december", "day":"Sunday"}
            ]
        }

        HorizontalHeaderView {
            width: root.width
            height: root.height - date.height
            interactive: false
            implicitWidth: 102
            implicitHeight: root.height - date.height
            model: 7
            columnSpacing: 2
            ScrollBar.vertical: ScrollBar { active: true; visible: true}
            delegate: Rectangle {
                id: _delegate
                width: 100
                implicitWidth: 102
                height: root.height
                implicitHeight: root.height
                ListView {
                    id: dayList
                    interactive: false
                    anchors.fill: parent
                    model: 24
                    spacing: 2
                    delegate: Rectangle {
                        width: 100
                        height: 30
                        color: "red"
                    }
                }

                Repeater {
                    model: root.model.meetings
                    Rectangle {
                        color: "purple"
                        width: 90
                        height: {
                            const startTimeParts = modelData.startTime.split(":");
                            const startTimeHour = startTimeParts[0]
                            const startTimeMinutes = startTimeParts[1]

                            const endTimeParts = modelData.endTime.split(":");
                            const endTimeHour = endTimeParts[0]
                            const endTimeMinutes = endTimeParts[1]

                            const hoursDiff = parseInt(endTimeHour) - parseInt(startTimeHour)
                            const minuteHeight = 30 / 30 //dayList.delegate.height
                            const hourHeight = 30 * 2
                            const finalHeight = hoursDiff * hourHeight + Math.abs(startTimeMinutes - endTimeMinutes) * minuteHeight
                            console.log("Calculate height", finalHeight)
                            return finalHeight
                        }
                        x: 0
                        y: {
                            const parts = modelData.startTime.split(":");
                            const hour = parts[0]
                            const minutes = parts[1]
                            const minuteHeight = 30 / 30
                            const hourHeight = 30 * 2
                            const finalY = parseInt(hour) * hourHeight + parseInt(minutes) * minuteHeight + parseInt(hour) * 4 //4 is 2 spacing + 2 half of hour
                            console.log("Calculate y", finalY)
                            return finalY
                        }
                    }
                    Component.onCompleted: {
                        console.log(_delegate.x, _delegate.y)
                    }
                }
            }
        }
    }

 }
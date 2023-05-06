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

        TableView {
            id: tableView
            columnSpacing: 1
            rowSpacing: 1
            width: root.width
            height: root.height- 50
            model: root.model
            clip: true

            delegate: Rectangle {
                id: delegate
                implicitWidth: 100
                implicitHeight: 60
                color: "green"

                Rectangle {
                    width: 90
                    height: {
                        if(display === undefined) return 0
                        const startTimeParts = display.startTime.split(":");
                        const startTimeHour = startTimeParts[0]
                        const startTimeMinutes = startTimeParts[1]

                        const endTimeParts = display.endTime.split(":");
                        const endTimeHour = endTimeParts[0]
                        const endTimeMinutes = endTimeParts[1]

                        const hoursDiff = parseInt(endTimeHour) - parseInt(startTimeHour)
                        const minuteHeight = delegate.implicitHeight / 30
                        const hourHeight = delegate.implicitHeight * 2
                        const finalHeight = hoursDiff * hourHeight + Math.abs(startTimeMinutes - endTimeMinutes) * minuteHeight
                        console.log("Calculate height", finalHeight)
                        return finalHeight
                    }
                    color: "purple"
                    anchors.top: delegate.top
                    anchors.topMargin: {
                        if(display === undefined) return 0

                        const parts = display.startTime.split(":")
                        const minutes = parts[1]
                        const minuteHeight = delegate.implicitHeight / 30
                        const finalMargin = minuteHeight * Math.abs(minutes - (row % 2 === 0 ? 0 : 30))
                        console.log("Calculate margin", finalMargin)
                    }
                    visible: display !== undefined && display.title !== ""
                }
            }
        }
    }

 }
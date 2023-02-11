import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root
    property var model

    Column {
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
                    interactive: false
                    anchors.fill: parent
                    model: 12
                    spacing: 2
                    delegate: Rectangle {
                        width: 100
                        height: 50
                        color: "red"
                    }
                }

                Repeater {
                    model: 2
                    Rectangle {
                        color: "purple"
                        width: 90
                        height: 30
                        x: 0
                        y: index === 0 ? 20 : 60
                    }
                    Component.onCompleted: {
                        console.log(_delegate.x, _delegate.y)
                    }
                }
            }
        }
    }

 }
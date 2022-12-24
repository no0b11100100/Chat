import QtQuick 2.0

 Rectangle {
    id: root
    property var model

    // Column {
        // spacing: 2
        Date {
            id: date
            width: root.width
            // anchors.right: root.right
            anchors.leftMargin: time.width
            anchors.left: root.left
            height: 40
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


        // Row {
            Time{
                id: time
                height: root.height - date.height
                anchors.right: main.left
                anchors.bottom: root.bottom
                width: 50
                model: [
                    {"time":"7:00"},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":"8:00"},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":"9:00"},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":""},
                    {"time":"10:00"},
                ]
            }

            ListView {
                id: main
                anchors.top: date.bottom
                // anchors.fill: parent
                width: root.width - time.width
                height: root.height - date.height
                anchors.right: root.right
                model: root.model
                spacing: 2
                orientation: ListView.Horizontal
                delegate: Rectangle{
                    id: _delegate
                    width: 100
                    height: root.height - date.height

                    border.color: "black"
                    border.width: 1

                    ListView{
                        id: meets
                        model: 15//[1,2,3,4,5,6,7,8,9]
                        anchors.fill: parent
                        spacing: 2
                        delegate: Rectangle{
                            width: meets.width
                            height: 40
                            border.color: "black"
                            border.width: 1
                            color: index % 2 === 0 ? "lightgrey": "lightblue"
                        }
                    }

                    Meetings {
                        model: [1,2]
                        coordinates: {"x":meets.x, "y":meets.y}
                    }

                }
            }
        // }
    // }

 }
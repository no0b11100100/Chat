import QtQuick 2.0
import QtQuick.Controls

Rectangle {
    id: root
    property var action
    property var router
    color: "green"

    Column {
        anchors.centerIn: parent
        spacing: 5

        Repeater {
            id: repeater
            model: ["Full name", "Enter email", "Enter password", "Confirm password"]
            TextAreaInput {
                width: 300
                height: 30
                placeholderText: modelData
            }
        }

        Button {
            text: "Sign Up"
            font.pointSize: 12
            width: 300

            onPressed: {
                console.log("Pressed on Sign Up button")

                var args = []
                for(var i = 0; i < repeater.count; i++) {
                    args.push(repeater.itemAt(i).text)
                }
                console.log("args ", args)
                root.action(args)
            }
        }

    } //Column

    RouteArea {
        anchors.bottom: root.bottom
        anchors.horizontalCenter: root.horizontalCenter
        anchors.bottomMargin: 20

        route: root.router
        labelText:  "Have an account?"
        linkText: "Sign In"
    }

} //Rectangle

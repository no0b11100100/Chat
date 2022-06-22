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
            model: ["Username or email", "Password"]
            TextAreaInput {
                width: 300
                height: 30
                placeholderText: modelData
            }
        }

        Button {
            text: "Sign In"
            font.pointSize: 12
            width: 300

            onPressed: {
                var args = []
                args.push(repeater.itemAt(0).text)
                args.push(repeater.itemAt(1).text)
                root.action(args)
            }
        }

    } //Column


    RouteArea {
        anchors.bottom: root.bottom
        anchors.horizontalCenter: root.horizontalCenter
        anchors.bottomMargin: 20

        route: root.router
        labelText: "Don't have an account?"
        linkText: "Sign Up"
    }

} //Rectangle

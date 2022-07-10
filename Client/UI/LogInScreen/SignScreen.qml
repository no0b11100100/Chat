import QtQuick 2.0
import QtQuick.Controls

Rectangle {
    id: root

    property string resultStatus: ""
    property var action
    property var router
    property string buttonText
    property string labelText: ""
    property string linkText: ""
    property var fields: []

    color: "green"

    Column {
        anchors.centerIn: parent
        spacing: 5

        Rectangle {
            width: 300
            height: 30
            color: "transparent"
            Text {
                anchors.horizontalCenter: parent.horizontalCenter
                text: root.resultStatus
                font.pointSize: 12
            }
        }

        Repeater {
            id: repeater
            model: root.fields
            TextAreaInput {
                width: 300
                height: 30
                placeholderText: modelData.text
                hideInput: modelData.hide
            }
        }

        Button {
            text: root.buttonText
            font.pointSize: 12
            width: 300

            onPressed: {
                var args = []
                for(var i = 0; i < repeater.count; i++) {
                    args.push(repeater.itemAt(i).text)
                }
                console.log("args ", args)
                root.action(...args)
            }
        }

    } //Column


    RouteArea {
        anchors.bottom: root.bottom
        anchors.horizontalCenter: root.horizontalCenter
        anchors.bottomMargin: 20

        route: root.router
        labelText: root.labelText
        linkText: root.linkText
    }
}

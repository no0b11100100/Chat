import QtQuick 2.0
import QtQuick.Controls

Row {
    id: root
    property var route
    property string labelText: ""
    property string linkText: ""

    spacing: 5
    Label {
        text: root.labelText
        font.pointSize: 12
        color: "black"
    }

    Text {
        id: text
        text: root.linkText
        font.pointSize: 12
        font.bold: true

        MouseArea {
            anchors.fill: parent
            hoverEnabled: true

            onClicked: {
                console.log("Clicked on Sign In text")
                root.route()
            }

            onEntered: {
                console.log("Curson over Sig In text")
                text.font.underline = true
            }

            onExited: {
                console.log("Curson leave Sig In text area")
                text.font.underline = false
            }
        } //MouseArea

    } //Text

} //Row

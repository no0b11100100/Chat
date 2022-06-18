import QtQuick 2.0
import QtQuick.Controls

Rectangle {
    id: _root
    property var action
    color: "green"

    Column {
        id: _column
        readonly property string emailText: email.text
        readonly property string passwordText: password.text
        anchors.centerIn: parent
        spacing: 5

        TextField{
            id: email
            readonly property string defaultPlaceholderText: "Username or email"
            placeholderText: email.defaultPlaceholderText
            placeholderTextColor: "grey"
            width: 300
            color: "black"
            font.pointSize: 16
            Component.onCompleted: {
                email.background.color = "white"
                email.background.opacity = 0.7
            }

            onPressed: {
                password.placeholderText = password.defaultPlaceholderText
                email.placeholderText = ""
            }
        }

        TextField{
            id: password
            readonly property string defaultPlaceholderText: "Password"
            placeholderText: password.defaultPlaceholderText
            placeholderTextColor: "grey"
            width: 300
            color: "black"
            font.pointSize: 16
            Component.onCompleted: {
                password.background.color = "white"
                password.background.opacity = 0.7
            }

            onPressed: {
                email.placeholderText = email.defaultPlaceholderText
                password.placeholderText = ""
            }
        }

    } //Column

    Button {
        text: "Sign In"
        anchors.top: _column.bottom
        anchors.topMargin: 20
        anchors.horizontalCenter: _column.horizontalCenter

        onPressed: {
            _root.action(_column.emailText, _column.passwordText)
        }
    }


} //Rectangle

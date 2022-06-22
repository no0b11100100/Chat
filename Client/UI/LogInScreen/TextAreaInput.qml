import QtQuick 2.0
import QtQuick.Controls

Rectangle{
    id: root
    property string placeholderText: ""
    readonly property string text: textInput.text

    TextField {
        id: textInput
        placeholderText: root.placeholderText
        placeholderTextColor: "grey"
        width: root.width
        height: root.height
        color: "black"
        font.pointSize: 16
        Component.onCompleted: {
            textInput.background.color = "white"
            textInput.background.opacity = 0.7
        }

        onActiveFocusChanged: {
            if (focus) {
                textInput.placeholderText = ""
            } else {
                textInput.placeholderText = root.placeholderText
            }
        }
    } //TextField

} //Rectangle

import QtQuick 2.0
import QtQuick.Controls

Rectangle {
    id: root

    height: input.height

    TextField {
        id: input
        placeholderText: "Enter the message"
        width: root.width
    }
}
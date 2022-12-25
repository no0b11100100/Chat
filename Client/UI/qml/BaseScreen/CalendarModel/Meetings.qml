import QtQuick 2.0

Rectangle {
    id: root

    property var model
    property var coordinates

    Repeater {
        model: root.model
        Rectangle {
            width: 85
            height: index === 0 ? 60 : 40
            color: "purple"
            y: root.coordinates.y + index === 0 ? 100 : 400
            x: root.coordinates.x
        }
    }

}
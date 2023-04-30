import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root

    property var model

    Row {
        Lists {
            id: lists
            width: 100
            height: root.height
            model: root.model.lists
        }
        Column {
            Repeater {
                model: root.model.tasks
                Task {
                    width: root.width - lists.width
                    height: 50
                }
            }
        }
    }
}

import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root

    property var model

    Row {
        ListView {
            id: lists
            height: root.height
            width: 100
            model: root.model.listsModel
            delegate: Rectangle {
                width: 98
                height: 30
                radius: 2
                border.color: "black"
                border.width: 1
                Text {
                    text: display.title
                    anchors.centerIn: parent
                }

                MouseArea {
                    anchors.fill: parent
                    onClicked: { root.model.listsModel.selectList(display.id) }
                }
            }
        }

        ListView {
            id: tasks
            height: root.height
            width: root.width - lists.width
            model: root.model.tasksModel

            delegate: Rectangle {}
        }
    }
}

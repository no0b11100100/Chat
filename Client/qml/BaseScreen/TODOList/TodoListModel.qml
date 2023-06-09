import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root

    property var model

    Row {
        Column {
            ListView {
                id: lists
                height: root.height - newListButton.height
                width: 200
                model: root.model.listsModel
                delegate: Rectangle {
                    width: parent.width
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

            TextField {
                id: newListButton
                width: lists.width
                height: 30
                placeholderText: "Add list"
                placeholderTextColor: "black"

                Keys.onReturnPressed: {
                    root.model.listsModel.addList(newListButton.text)
                    newListButton.text = ""
                }
                Keys.onEnterPressed: {
                    root.model.listsModel.addList(newListButton.text)
                    newListButton.text = ""
                }
            }

        }

        Column {
            ListView {
                id: tasks
                height: root.height - newListTask.height
                width: root.width - lists.width
                model: root.model.tasksModel

                delegate: Rectangle {
                    id: delegate
                    width: parent.width - 10
                    height: 30
                    border.color: "black"
                    border.width: 1
                    radius: 5
                    CheckBox {
                        checked: false
                        text: "|" + display.title + "|"
                        height: delegate.height - 10

                        onCheckedChanged: {
                            console.log("Status changed", checked)
                            root.model.tasksModel.setTaskState("", checked)
                        }
                    }
                }
            }

            TextField {
                id: newListTask
                width: tasks.width
                height: 30
                placeholderText: "Add task"
                placeholderTextColor: "black"

                Keys.onReturnPressed: {
                    root.model.tasksModel.addTask(newListTask.text)
                    newListTask.text = ""
                }
                Keys.onEnterPressed: {
                    root.model.tasksModel.addTask(newListTask.text)
                    newListTask.text = ""
                }
            }
        }
    }
}

import QtQuick 2.0

Rectangle {
    id: root
    property var model

    ListView {
        anchors.fill: parent
        model: root.model
        spacing: 2
        delegate: Notification{
            width: root.width
            height: 40 //TODO
            model: display
        }

        Component.onCompleted: {
            console.log(root.model === undefined, root.model.name)
        }
    }
}

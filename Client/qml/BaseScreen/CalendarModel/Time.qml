import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root
    property var model

    VerticalHeaderView {
        anchors.fill: parent
        interactive: false
        model: root.model

        delegate: Rectangle{
            implicitHeight:42
            width: 50
            Text {
                anchors.centerIn: parent
                text: root.model[index].time
            }
        }
    }
}

import QtQuick 2.15
import QtQuick.Controls 2.15

Rectangle {
    id: root
    property var model

    HorizontalHeaderView {
        implicitWidth: 100
        anchors.fill: parent
        interactive: false
        rowSpacing: 10
        model: root.model

        delegate: Rectangle{
            implicitWidth: 102
            height: 40
            Column {
                // width: 102
                anchors.fill: parent
                Text {
                    text: root.model[index].date
                }
                Text {
                    text: root.model[index].day
                }
            }
        }
    }
}

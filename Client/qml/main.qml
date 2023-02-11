import QtQuick 2.14
import QtQuick.Window 2.14
import QtQuick.Controls 2.14

import Models 1.0


Window {
    id: _root
    width: 900
    height: 600
    visible: true
    title: qsTr("Hello World")
    color: "green"

    ScreenController {
        id: screenConttroller
        anchors.fill: parent
        model: _models.model
    }

    Backend {
        id: _models

        onModelChanged: {
            screenConttroller.changeScreen()
        }
    }

//     ScrollView {
//         width: 900
//         height: _root.height/ 3
//         contentHeight: height * 2

//     Row {
//         spacing: 2
//         // anchors.fill: parent
//         Repeater {
//             // anchors.fill: parent
//             model: 7
//             Rectangle {
//                 width: 100
//                 height: _root.height
//                 ListView {
//                     id: listView
//                     anchors.fill: parent
//                     spacing: 2
//                     model: 5
//                     delegate: Rectangle {
//                         width: 100
//                         height: 50
//                         color: "red"
//                     }
//                 }
//                 Repeater {
//                     model: 2
//                     Rectangle {
//                         color: "purple"
//                         width: 90
//                         height: 30
//                         x: listView.x
//                         y: index === 0 ? listView.y + 20 : listView.y + 100
//                     }
//                 }
//             }
//         }
//     }
// }

}
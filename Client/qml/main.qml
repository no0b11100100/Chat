import QtQuick 2.0
import QtQuick.Window 2.0
import QtQuick.Controls 2.0

// import Models 1.0

Window {
    id: _root
    width: 640
    height: 480
    visible: true
    title: qsTr("Hello World")
    color: "green"

    // ScreenController {
    //     id: screenConttroller
    //     anchors.fill: parent
    //     model: _models.model
    // }

    // Backend {
    //     id: _models

    //     onModelChanged: {
    //         screenConttroller.changeScreen()
    //     }
    // }

}

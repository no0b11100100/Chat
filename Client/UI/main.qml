import QtQuick
import QtQuick.Window
import QtQuick.Controls

import "LogInScreen"

import Models 1.0

Window {
    id: _root
    width: 640
    height: 480
    visible: true
    title: qsTr("Hello World")
    color: "green"

    LogInScreen {
        action: _models.model.triggerLogIn
        anchors.fill: parent
    }

    Backend {
        id: _models
    }

}

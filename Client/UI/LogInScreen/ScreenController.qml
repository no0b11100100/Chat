import QtQuick 2.0

Rectangle {
    id: root
    property var model

    Loader {
        id: screenLoader
        anchors.fill: parent
    }

    Component {
        id: signIn

        SignIn {
            anchors.fill: parent
            action: root.model.signIn
            router: function() { screenLoader.sourceComponent = signUp }
        }
    }

    Component {
        id: signUp

        SignUp {
            anchors.fill: parent
            action: root.model.signUp
            router: function() { screenLoader.sourceComponent = signIn }
        }
    }

    Component.onCompleted: {
        screenLoader.sourceComponent = signIn
    }
}
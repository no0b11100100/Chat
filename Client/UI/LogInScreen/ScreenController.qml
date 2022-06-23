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

        SignScreen {
            anchors.fill: parent
            action: root.model.signIn
            buttonText: "Sign In"
            router: function() { screenLoader.sourceComponent = signUp }
            labelText: "Don't have an account?"
            linkText: "Sign Up"
            fields: ["Username or email", "Password"]
        }
    }

    Component {
        id: signUp

        SignScreen {
            anchors.fill: parent
            action: root.model.signUp
            buttonText: "Sign Up"
            router: function() { screenLoader.sourceComponent = signIn }
            labelText: "Have an account?"
            linkText: "Sign In"
            fields: ["Full name", "Enter email", "Enter password", "Confirm password"]
        }
    }

    Component.onCompleted: {
        screenLoader.sourceComponent = signIn
    }
}

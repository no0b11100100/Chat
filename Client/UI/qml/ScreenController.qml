import QtQuick 2.0
import QtQml

import "LogInScreen"
import "BaseScreen"

Rectangle {
    id: root
    property var model
    property var changeScreen: function() {
        console.log("Change screen")
        screenLoader.sourceComponent = baseScreen
    }

    Loader {
        id: screenLoader
        anchors.fill: parent
    }

    Component {
        id: signIn

        SignScreen {
            id: signInScreen
            anchors.fill: parent
            action: root.model.signIn
            buttonText: "Sign In"
            router: function() { screenLoader.sourceComponent = signUp }
            labelText: "Don't have an account?"
            linkText: "Sign Up"
            fields: [
                {text:"Username or email", hide: false},
                {text:"Password", hide: true}
            ]


            Connections {
                target: root.model
                function onStatusMessage(message) { signInScreen.resultStatus = message }
            }
        }
    }

    Component {
        id: signUp

        SignScreen {
            id: signUpScreen
            anchors.fill: parent
            action: root.model.signUp
            buttonText: "Sign Up"
            router: function() { screenLoader.sourceComponent = signIn }
            labelText: "Have an account?"
            linkText: "Sign In"
            fields: [
                {text:"Full name", hide: false},
                {text:"Nickname", hide: false},
                {text:"Enter email", hide: false},
                {text:"Enter password", hide: true},
                {text:"Confirm password", hide: true},
            ]

            Connections {
                target: root.model
                function onStatusMessage(message) { signUpScreen.resultStatus = message }
            }
        }
    }

    Component {
        id: baseScreen

        BaseScreen {
            anchors.fill: parent
            model: root.model
            Component.onCompleted: {
                console.log("Create BaseScreen", root.model === undefined)
            }
        }
    }

    Component.onCompleted: {
        screenLoader.sourceComponent = signIn
    }

}

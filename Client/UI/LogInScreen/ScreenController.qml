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

            Component.onCompleted: {
                root.model.statusMessage.connect(function(message) { signInScreen.resultStatus = message })
            }
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
            fields: [
                {text:"Full name", hide: false},
                {text:"Enter email", hide: false},
                {text:"Enter password", hide: true},
                {text:"Confirm password", hide: true},
            ]
        }
    }

    Component.onCompleted: {
        screenLoader.sourceComponent = signIn
    }

}

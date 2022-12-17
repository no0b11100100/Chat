import QtQuick 2.0

import "ChatListModel"
import "ChatModel"
import "NotificationModel"

Rectangle {
    id: root
    property var model

    Row {
        id: row
        // property bool isNotificationsOpened: true//false
        spacing: 1
        SideBar {
            id: sidebar
            width: 30
            height: root.height
            buttons: [{"id": "notificationIDButton", "text":"notification"}, {"id": "chatIDButton", "text":"chat"}, {"id": "textIDButton", "text":"mail"}, {"id": "todoIDButton", "text":"TODO"}]
            action: function(id) {
                console.log("Press sidebar", id)
                if(id === "notificationIDButton")
                    pageLoader.sourceComponent = notification
                if(id === "chatIDButton")
                    pageLoader.sourceComponent = chats
            }
        }

        Loader {
            id: pageLoader
            width: root.width / 4 - sidebar.width
            height: root.height
        }

        Component {
            id: chats
            ChatListModel {
                id: _chats
                height: root.height
                width: root.width / 4 - sidebar.width
                model: root.model.chatListModel
                // visible: !row.isNotificationsOpened
                Component.onCompleted: {
                    console.log("Create ChatListModel", root.model.chatListModel === undefined, _chats.model.name)
                }
            }
        }

        Component {
            id: notification
            NotificationModel {
                id: _notification
                height: root.height
                width: chats.width
                model: root.model.notificationListModel
                // visible: row.isNotificationsOpened
                Component.onCompleted: {
                    console.log("Create NotificationModel", root.model.notificationListModel === undefined, _notification.model.name)
                }
            }
        }

        ChatModel {
            id: chatMessages
            height: root.height
            width: root.width - pageLoader.width - sidebar.width
            model: root.model.chatModel
            Component.onCompleted: {
                console.log("Create ChatModel", root.model.chatModel === undefined, chatMessages.model.name)
            }
        }

        Component.onCompleted: {
            pageLoader.sourceComponent = chats
        }
    }
}

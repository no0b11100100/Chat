#include <QGuiApplication>
#include <QQmlApplicationEngine>

#include "src/App.hpp"

// #include "communication/json/Value/Value.h"

// enum class ResponseStatus {
//     OK = 0,
//     ERROR = 1,
//     CHECK=2
// };

// struct UserInfo : public Types::ClassParser {
//   std::string UserID;
//   std::string Name;
//   std::string NickName;
//   std::string Photo;
//   std::vector<std::string> Chats;
//   std::string Email;
//   std::string Password;
//   virtual json toJson() const override {
//     json js({});
//     js["UserID"] = UserID;
//     js["Name"] = Name;
//     js["NickName"] = NickName;
//     js["Photo"] = Photo;
//     js["Chats"] = Chats;
//     js["Email"] = Email;
//     js["Password"] = Password;

//     // std::cout << "UserInfo\n" << js.dump() << "\n" << js["Chats"].isVector() << std::endl;

//     return js;
//   }

//   virtual void fromJson(json js) override {
//     UserID = static_cast<std::string>(js["UserID"]);
//     Name = static_cast<std::string>(js["Name"]);
//     NickName = static_cast<std::string>(js["NickName"]);
//     Photo = static_cast<std::string>(js["Photo"]);
//     if(js["Chats"].isNull()) Chats = std::vector<std::string>();
//     else Chats = static_cast<std::vector<std::string>>(js["Chats"]);
//     Email = static_cast<std::string>(js["Email"]);
//     Password = static_cast<std::string>(js["Password"]);
//   }
// };

// struct Response : public Types::ClassParser {
//   UserInfo Info;
//   ResponseStatus Status;
//   std::string StatusMessage;
//   virtual json toJson() const override {
//     json js({});
//     js["Info"] = Info;
//     js["Status"] = Status;
//     js["StatusMessage"] = StatusMessage;
//     return js;
//   }

//   virtual void fromJson(json js) override {
//     Info = static_cast<UserInfo>(js["Info"]);
//     Status = static_cast<ResponseStatus>(js["Status"]);
//     StatusMessage = static_cast<std::string>(js["StatusMessage"]);
//   }
// };

int main(int argc, char *argv[])
{
#if QT_VERSION < QT_VERSION_CHECK(6, 0, 0)
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif

    QGuiApplication app(argc, argv);

    QQmlApplicationEngine engine;

    qmlRegisterType<App, 1>("Models", 1, 0, "Backend");
    qRegisterMetaType<user::Response>("user::Response");

    // used for qml logs
    qSetMessagePattern("%{time HH:mm:ss }%{file}:%{line}: %{message}");

    const QUrl url(QStringLiteral("qrc:/qml/main.qml"));
    QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
                     &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
            QCoreApplication::exit(-1);
    }, Qt::QueuedConnection);
    engine.load(url);


    // MessageData data;
    // data.Type = EnumClass::ERROR;
    // data.Payload = json::array({data});
    // data.Topic = "12345";
    // json js = data.toJson();
    // std::cout << js.dump() << std::endl;

    // js["Topic"] = "1";
    // data = js;

    // std::cout << data.Topic << std::endl;

    // std::string value = "{\"Endpoint\":\"UserService.SignUp\",\"Topic\":\"1673830052388\",\"Payload\":{\"Info\":{\"UserID\":\"b7e7e0b5-31bc-4453-8edd-6e53a23d3ac7\",\"Name\":\"\",\"NickName\":\"\",\"Photo\":\"\",\"Chats\":null,\"Email\":\"\",\"Password\":\"\"},\"Status\":0,\"StatusMessage\":\"\"},\"Type\":-1889978384}";

    // json jj = json::parse(value);

    // std::cout << jj << std::endl;

    // MessageData msg;
    // msg = jj;
    // std::cout << msg.Topic << std::endl;


    // std::vector<std::string> v;
    // v.push_back("1");
    // v.push_back("2");
    // v.push_back("3");

    // std::string n = R"({"Endpoint":"UserService.SignUp","Topic":"1673906710152","Payload":{"Info":{"UserID":"c7cc6705-c0ff-4ef9-9897-e4e501715b86","Name":"","NickName":"","Photo":"","Chats":null,"Email":"","Password":""},"Status":0,"StatusMessage":""},"Type":-2127626304})";

    // json js = json::parse(n);
    // // js["Vector"] = v;
    // // js = Response();

    // std::cout << js << std::endl;

    // Response r;
    // r = js["Payload"];


    return app.exec();
}

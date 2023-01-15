#include <QGuiApplication>
#include <QQmlApplicationEngine>

#include "src/App.hpp"

// #include "communication/json/Value/Value.h"

enum class EnumClass {
    OK = 0,
    ERROR = 1,
    CHECK=2
};

struct MessageData : public Types::ClassParser {
  std::string Endpoint;
  std::string Topic;
  json Payload;
  EnumClass Type;
  virtual json toJson() const override {
    json js({});
    js["Endpoint"] = Endpoint;
    js["Topic"] = Topic;
    js["Payload"] = Payload;
    js["Type"] = Type;
    return js;
  }

  virtual void fromJson(json js) override {
    Endpoint = static_cast<std::string>(js["Endpoint"]);
    Topic = static_cast<std::string>(js["Topic"]);
    Payload = static_cast<std::string>(js["Payload"]);
    Type = static_cast<EnumClass>(js["Type"]);
  }
};


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
    // std::cout << data.toJson().dump() << std::endl;


    return app.exec();
}

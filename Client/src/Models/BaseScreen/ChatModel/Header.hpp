#pragma once

#include <QObject>
#include <QDebug>

class Header : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title NOTIFY titleChanged)
    Q_PROPERTY(QString secondLine READ secondLine NOTIFY secondLineChanged)
public:
    Header(QObject* parent = nullptr)
        : QObject{parent}
    {}

    void SetTitle(const QString& title)
    {
        m_title = title;
        emit titleChanged();
    }

    void SetSecondLine(const QString& secondLine)
    {
        m_secondLine = secondLine;
        emit secondLineChanged();
    }

    void SetCallAction(std::function<void()> action) { m_callAction = action; }

    QString title() const { return m_title; }
    QString secondLine() const { return m_secondLine; }

    Q_INVOKABLE void call() {
        qDebug() << "Press call in chat";
        m_callAction();
    }

signals:
    void titleChanged();
    void secondLineChanged();

private:
    QString m_title;
    QString m_secondLine;
    std::function<void()> m_callAction;
};
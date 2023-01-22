#pragma once
#include <QAudioInput>
#include <QAudioOutput>
#include <QDebug>
#include <QTimer>
#include <memory>

class Audio : public QObject
{
    std::unique_ptr<QAudioInput> m_input;
    std::unique_ptr<QAudioOutput> m_output;
    std::unique_ptr<QIODevice> m_inputStream;
    std::unique_ptr<QIODevice> m_outputStream;
    QTimer m_timer;
    std::vector<std::function<void(QByteArray)>> m_inputObservers;

public:
    Audio(QObject* parent=nullptr)
    : m_input{nullptr},
    m_output{nullptr},
    m_inputStream{nullptr},
    m_outputStream{nullptr}
    {
        QAudioFormat format;
        // Set up the desired format, for example:
        format.setSampleRate(8000);
        format.setChannelCount(1);
        format.setSampleSize(8);
        format.setCodec("audio/pcm");
        format.setByteOrder(QAudioFormat::LittleEndian);
        format.setSampleType(QAudioFormat::UnSignedInt);

        QAudioDeviceInfo info = QAudioDeviceInfo::defaultInputDevice();
        if (!info.isFormatSupported(format)) {
            qWarning() << "Default format not supported, trying to use the nearest.";
            format = info.nearestFormat(format);
        }

        m_input.reset(new QAudioInput(format, this));
        m_output.reset(new QAudioOutput(format, this));
        // m_input->setNotifyInterval(3000);

        // m_stream = std::move(std::unique_ptr<QIODevice>(m_input->start())); //Works

        // connect(m_input.get(), &QAudioInput::notify, this, &Audio::handleNotify);
        // connect(m_input.get(), &QAudioInput::stateChanged, this, &Audio::handleStateChanged);
        // connect(m_stream.get(), &QIODevice::readyRead, this, &Audio::readInput); //Works

        // QTimer::singleShot(3000, this, &Audio::stopRecording);
        // m_stream = std::move(std::unique_ptr<QIODevice>(m_input->start()));

        // qDebug() << "Interval" << m_input->notifyInterval();

        // m_output.reset(new QAudioOutput(format, this));
        // connect(m_output.get(), &QAudioOutput::stateChanged, this, &Audio::handleOutputStateChanged);

        // m_timer.setInterval(3000);
        // m_timer.callOnTimeout(this, &Audio::handleInterval);
        // m_timer.start();
    }

    void startStream()
    {
        m_inputStream = std::move(std::unique_ptr<QIODevice>(m_input->start()));
        m_outputStream = std::move(std::unique_ptr<QIODevice>(m_output->start()));
        connect(m_inputStream.get(), &QIODevice::readyRead, this, &Audio::readInput);
    }

    void SubscribeOnAudioInput(std::function<void(QByteArray)> callback)
    {
        m_inputObservers.push_back(callback);
    }

public slots:
    void receiveStream(QByteArray data)
    {
        qDebug() << "Output <" << data << ">";
        m_outputStream->write(data);
    }

    void readInput()
    {
        auto data = m_inputStream->readAll();
        qDebug() << "Input <" << data << ">";
        for(auto callback : m_inputObservers)
            callback(data);
    }

    // void handleNotify()
    // {
    //     qDebug() << "Handle notify" << m_stream->size();
    // }

    // void handleInterval()
    // {
    //     qDebug() << "Handle interval";
    // }

    // void stopRecording()
    // {
    //     m_output->start(m_stream.get());
    //     // m_input->stop();
    // }

    // void handleStateChanged(QAudio::State newState)
    // {
    //     switch (newState) {
    //         case QAudio::StoppedState:
    //             if (m_input->error() != QAudio::NoError) {
    //                 // Error handling
    //             } else {
    //                 // Finished recording
    //             }
    //             break;

    //         case QAudio::ActiveState:
    //             // Started recording - read from IO device
    //             break;

    //         default:
    //             // ... other cases as appropriate
    //             break;
    //     }
    // }

    // void handleOutputStateChanged(QAudio::State newState)
    // {
    //     switch (newState) {
    //         case QAudio::IdleState:
    //             // Finished playing (no more data)
    //             m_output->stop();
    //             m_input->stop();
    //             break;

    //         case QAudio::StoppedState:
    //             // Stopped for other reasons
    //             if (m_output->error() != QAudio::NoError) {
    //                 // Error handling
    //             }
    //             break;

    //         default:
    //             // ... other cases as appropriate
    //             break;
    //     }
    // }

};

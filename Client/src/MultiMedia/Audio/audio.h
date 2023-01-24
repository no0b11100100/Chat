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
        // QAudioFormat inputFormat;
        // // Set up the desired format, for example:
        // inputFormat.setSampleRate(8000);
        // inputFormat.setChannelCount(1);
        // inputFormat.setSampleSize(8);
        // inputFormat.setCodec("audio/pcm");
        // inputFormat.setByteOrder(QAudioFormat::LittleEndian);
        // inputFormat.setSampleType(QAudioFormat::UnSignedInt);

        QAudioDeviceInfo inputInfo = QAudioDeviceInfo::defaultInputDevice();
        QAudioFormat inputFormat = inputInfo.preferredFormat();
        // if (!inputInfo.isFormatSupported(inputFormat)) {
        //     qWarning() << "Default format not supported, trying to use the nearest.";
        //     inputFormat = inputInfo.nearestFormat(inputFormat);
        // }

        // QAudioFormat outputFormat;
        // // Set up the format, eg.
        // outputFormat.setSampleRate(8000);
        // outputFormat.setChannelCount(1);
        // outputFormat.setSampleSize(8);
        // outputFormat.setCodec("audio/pcm");
        // outputFormat.setByteOrder(QAudioFormat::LittleEndian);
        // outputFormat.setSampleType(QAudioFormat::UnSignedInt);

        QAudioDeviceInfo outputInfo(QAudioDeviceInfo::defaultOutputDevice());
        QAudioFormat outputFormat = outputInfo.preferredFormat();

        // if (!outputInfo.isFormatSupported(outputFormat)) {
        //     qWarning() << "Raw audio format not supported by backend, cannot play audio.";
        //     return;
        // }

        qDebug() << "Output format" << outputFormat.sampleRate() << outputFormat.channelCount();
        qDebug() << "Input format" << inputFormat.sampleRate() << inputFormat.channelCount();

        qDebug() << "Input device" << inputInfo.deviceName() << "Output device" << outputInfo.deviceName();

        for(auto d : QAudioDeviceInfo::availableDevices(QAudio::AudioInput))
            qDebug() << "Input" << d.deviceName();
        for(auto d : QAudioDeviceInfo::availableDevices(QAudio::AudioOutput))
            qDebug() << "Output" << d.deviceName();

        m_input.reset(new QAudioInput(inputFormat, this));
        m_input->setBufferSize(4096);
        m_output.reset(new QAudioOutput(outputFormat, this));
        m_output->setBufferSize(4096);
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

        // SubscribeOnAudioInput([this](QByteArray data){receiveStream(data); });
    }

    void SubscribeOnAudioInput(std::function<void(QByteArray)> callback)
    {
        m_inputObservers.push_back(callback);
    }

public slots:
    void receiveStream(QByteArray data)
    {
        qDebug() << "Output <" << data.size() << ">";
        m_outputStream->write(QByteArray::fromHex(data));
    }

    void readInput()
    {
        // m_inputStream->startTransaction();
        auto data = m_inputStream->readAll();
        // m_inputStream->commitTransaction();
        qDebug() << "Input <" << data.toHex(0).size() << ">";
        for(auto callback : m_inputObservers)
            callback(data.toHex(0));
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

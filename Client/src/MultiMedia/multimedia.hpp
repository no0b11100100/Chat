#pragma once

#include "Audio/audio.h"

class Multimedia {
    Audio m_audio;
public:
    Multimedia() = default;

    Audio& audio() { return m_audio; }

};

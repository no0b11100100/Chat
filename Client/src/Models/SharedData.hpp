#pragma once

#include <string>
#include <functional>

class SharedData
{
    SharedData() = default;
    SharedData(SharedData&&) = delete;
    SharedData(const SharedData&) = delete;

    std::string m_emailField;

public:
    static SharedData& getConnector()
    {
        static SharedData instance;
        return instance;
    }

    void SaveEmailField(std::string email) { m_emailField = email; }
    std::string GeEmailField() const { return m_emailField; }
};

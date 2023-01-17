#pragma once
#include <limits>
#include <type_traits>
#include <iostream>

#include "../Base.hpp"

namespace Types {

class String : public Base
{
    std::string m_value;

public:
    String();

    template<class T>
    requires StringValue<T>
    String(T&& value)
        : m_value{value}
    {}

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    String substr(size_t pos = 0, size_t len = std::string::npos) const;

    template<class T>
    requires StringValue<T>
    void operator=(T&& value) { m_value = value; }

    template<class T>
    requires StringValue<T>
    operator T() { return m_value; }

    friend std::ostream& operator <<(std::ostream & os, const String& i)
    {
        os << "\"" << i.m_value << "\"";
        return os;
    }

    friend std::istream& operator >>(std::istream & is, String& value)
    {
        is >> value.m_value;
        return is;
    }
};

}

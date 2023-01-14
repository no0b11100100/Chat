#pragma once
#include <limits>
#include <type_traits>
#include <iostream>

#include "../Base.hpp"

namespace Types {

class Number : public Base
{
    long double m_value;

public:
    Number();

    template<class T>
    requires NumberValue<T>
    Number(T&& value)
        : m_value{value}
    {}

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    template<class T>
    requires NumberValue<T>
    void operator=(T&& value) { m_value = value; }

    template<class T>
    requires NumberValue<T>
    operator T()
    {
        if(m_value > std::numeric_limits<T>::max()) return T();
        return static_cast<T>(m_value);
    }

    friend std::ostream& operator <<(std::ostream & os, const Number& i)
    {
        os << std::to_string(i.m_value);
        return os;
    }

    friend std::istream& operator >>(std::istream & is, Number& value)
    {
        is >> value.m_value;
        return is;
    }

};

}

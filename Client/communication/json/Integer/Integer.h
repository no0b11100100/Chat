#pragma once
#include <limits>
#include  <type_traits>
#include <iostream>

#include "../Base.hpp"

namespace Types {

class Integer : public Base
{

    long long m_value;

public:
    Integer();

    template<class T>
    requires IntegerValue<T>
    Integer(T&& value)
        : m_value{static_cast<long long>(value)}
    {}

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    template<class T>
    requires IntegerValue<T>
    void operator=(T&& value) { m_value = value; }

    template<class T>
    requires IntegerValue<T>
    operator T()
    {
        if(m_value > std::numeric_limits<T>::max()) return T();
        return static_cast<T>(m_value);
    }

    friend std::ostream& operator <<(std::ostream & os, const Integer& i)
    {
        os << std::to_string(i.m_value);
        return os;
    }

    friend std::istream& operator >>(std::istream & is, Integer& value)
    {
        is >> value.m_value;
        return is;
    }
};

}

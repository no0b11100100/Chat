#pragma once
#include <limits>
#include <type_traits>
#include <iostream>

#include "../Base.hpp"

namespace Types {

class Bool : public Base
{
    bool m_value;

public:
    Bool();

    template<class T>
    requires BoolValue<T>
    Bool(T&& value)
        : m_value{value}
    {}

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    template<class T>
    requires BoolValue<T>
    void operator=(T&& value) { m_value = value; }

    template<class T>
    requires BoolValue<T>
    operator T() { return m_value; }

    friend std::ostream &operator <<(std::ostream&, const Bool&);
    friend std::istream &operator >>(std::istream&, Bool&);
};

}

#pragma once
#include <limits>
#include <type_traits>
#include <iostream>

#include "../Base.hpp"

namespace Types {

class Null : public Base
{
public:
    Null() = default;

    template<class T>
    requires NullValue<T>
    Null(T&&)
    {}

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    template<class T>
    requires NullValue<T>
    void operator=(T&&) {}

    template<class T>
    requires NullValue<T>
    operator T() { return nullptr; }

    friend std::ostream& operator <<(std::ostream & os, const Null&)
    {
        os << "null";
        return os;
    }

};

}

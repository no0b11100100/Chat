#pragma once
#include <string>

template<class T>
struct Pair
{
    std::string key;
    T value;

    Pair(const std::string& _key, const T& _value)
        : key{_key},
          value{_value}
    {}
};

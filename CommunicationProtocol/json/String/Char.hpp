#pragma once
#include <string>
#include <iostream>
#include <functional>

#include "String.h"

// class Char : public Types::Base
// {
// public:
//     char value;

//     Types::Type type() const override { return Types::Type::STRING; }
//     bool isEmpty() const override {return true; }
//     std::size_t size() const override { return 0; }

//     operator char() const { return value; }
//     operator const char() const { return value; }

//     // friend std::ostream& operator <<(std::ostream& os, Char& p)
//     // {
//     //     os << p.value;
//     //     return os;
//     // }
// };

template<class ValueType>
class String
{
    std::basic_string<ValueType> m_value;
public:
    String(std::string str)
    // : m_value{str}
    {
        std::cout << "Create string\n";
        for (char c : str)
        {
            // Char ch;
            // ch.value = c;
            // m_value.push_back(ch);
            // m_value.push_back(ValueType::createFromChar(c));
        }
    }

    void print()
    {
        // for(auto v : m_value)
        //     std::cout << v;
    }

};

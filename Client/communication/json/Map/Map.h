#pragma once

#include "Pair.hpp"
#include "../Base.hpp"
#include "../Iterator/Iteartor.h"

#include <iostream>
#include <algorithm>

namespace Types {

template<class ValueType>
class Map : public Base
{
    std::vector<Pair<ValueType>> m_value;

public:
    Map() = default;

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    auto find(const std::string& key)
    {
        return std::find_if(m_value.begin(), m_value.end(), [&key](const Pair<ValueType>& p)
        {
            return p.key == key;
        });
    }

    bool contains(const std::string& key);

    template<class T>
    requires StringValue<T>
    ValueType& at(T&& key)
    {
        auto it = find(key);
        if(it == m_value.end())
        {
            m_value.emplace_back(key, ValueType());
            it = std::prev(m_value.end());
        }

        return it->value;
    }

    template<class T>
    requires StringValue<T>
    const ValueType& at(T&& key) const
    {
        auto it = find(key);
        if(it == m_value.end())
        {
            m_value.emplace_back(key, ValueType());
            it = std::prev(m_value.end());
        }

        return it->value;
    }

    auto begin() { return Iterator<ValueType>(m_value.begin(), m_value.end()); }
    auto end() { return Iterator<ValueType>(m_value.end(), m_value.end()); }

};

}

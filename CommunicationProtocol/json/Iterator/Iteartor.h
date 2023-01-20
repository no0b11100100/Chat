#pragma once
#include <map>
#include <deque>
#include <vector>
#include <variant>

#include "../Map/Pair.hpp"

#include <iostream>

template <class Value>
class Iterator
{
    using iterator_category = std::bidirectional_iterator_tag;
    using difference_type   = std::ptrdiff_t;
    using value_type        = Value;
    using pointer           = value_type*;
    using reference         = value_type&;

    template<class... Ts> struct overloaded : Ts... { using Ts::operator()...; };

    using VectorIterator = typename std::deque<Value>::iterator;
    using MapIterator = typename std::vector<Pair<Value>>::iterator;

    template<class Info>
    struct IteratorInfo
    {
        Info currentIterator;
        Info endIterator;
    };

    std::variant<IteratorInfo<MapIterator>, IteratorInfo<VectorIterator>> m_iterator;
    std::pair<Value, pointer> m_data;
public:
    Iterator()
        : m_data{nullptr, nullptr}
    {}

    Iterator(MapIterator it, MapIterator end)
        : m_iterator{IteratorInfo<MapIterator>{it, end}},
          m_data{it != end ? it->key : "", it != end ?  &it->value : nullptr}
    {}

    Iterator(std::size_t index, VectorIterator it, VectorIterator end)
        : m_iterator{IteratorInfo<VectorIterator>{it, end}},
          m_data{index, it != end ? it.operator->() : nullptr}
    {}

    Value key() const { return m_data.first; }
    Value value() const { return m_data.second; }
    reference value() { return *(m_data.second); }
    Iterator& next() { return operator++(); }
    Iterator& prev() { return operator--(); }

    reference operator*() const { return *(m_data.second); }
    pointer operator->() { return m_data.second; }

    // Prefix
    Iterator& operator++()
    {
        std::visit(overloaded{
                       [&](IteratorInfo<VectorIterator>& it)
                       {
//                           m_iterator = ++it;
//                           std::size_t index = m_data.first;
//                           m_data.first = index+1;
//                           m_data.second = it.operator->();

                           ++it.currentIterator;
                           if(it.currentIterator != it.endIterator)
                           {
                               std::size_t index = m_data.first;
                               m_data.first = index+1;
                               m_data.second = it.currentIterator.operator->();
                           }
                       },
                       [&](IteratorInfo<MapIterator>& it)
                       {
                           ++it.currentIterator;
                           if(it.currentIterator != it.endIterator)
                           {
                               m_data.first = it.currentIterator->key;
                               m_data.second = &it.currentIterator->value;
                           }
                       }
                   }, m_iterator);
        return *this;
    }

    // Postfix
    Iterator operator++(int)
    {
        Iterator tmp = *this;
        ++(*this);
        return tmp;
    }

    // Prefix
    Iterator& operator--()
    {
                                    std::visit(overloaded{
                                                   [&](IteratorInfo<VectorIterator>& it)
                                                   {
                                                       --it.currentIterator;
                                                       if(it.currentIterator != it.endIterator)
                                                       {
                                                           std::size_t index = m_data.first;
                                                           m_data.first = index-1;
                                                           m_data.second = &it.currentIterator.operator->();
                                                       }
                                                   },
                                                   [&](IteratorInfo<MapIterator>& it)
                                                   {
                                                       --it.currentIterator;
                                                       if(it.currentIterator != it.endIterator)
                                                       {
                                                           m_data.first = it.currentIterator->key;
                                                           m_data.second = &it.currentIterator->value;
                                                       }
                                                   }
                                               }, m_iterator);
        return *this;
    }

    // Postfix
    Iterator operator--(int)
    {
        Iterator tmp = *this;
        --(*this);
        return tmp;
    }

    friend bool operator== (const Iterator& lhs, const Iterator& rhs)
    {
        if (lhs.m_iterator.index() != rhs.m_iterator.index()) return false;
        bool result = false;

        std::visit(overloaded{
                       [&](IteratorInfo<VectorIterator> lhs, IteratorInfo<VectorIterator> rhs){ result = (lhs.currentIterator==rhs.currentIterator); },
                       [&](IteratorInfo<MapIterator> lhs, IteratorInfo<MapIterator> rhs){ result = (lhs.currentIterator==rhs.currentIterator); },
                       [&](IteratorInfo<VectorIterator>, IteratorInfo<MapIterator>){ result = false; },
                       [&](IteratorInfo<MapIterator>, IteratorInfo<VectorIterator>){ result = false; }
                   }, lhs.m_iterator, rhs.m_iterator);

        return result;
    }

    friend bool operator!= (const Iterator& lhs, const Iterator& rhs) { return !(lhs == rhs); };
};

#pragma once

//#include "../Value/Value.h"
#include "../Base.hpp"
#include "../Iterator/Iteartor.h"
#include <algorithm>

namespace Types {

template<class ValueType>
class Vector : public Base
{
    std::deque<ValueType> m_value;

    template<class T>
    void initFromContainer(const T& value)
    {
        if constexpr(is_specialization<std::decay_t<T>, std::stack>::value)
        {
            T tmpValue = value;
            while(!tmpValue.empty())
            {
                m_value.push_back(tmpValue.top());
                tmpValue.pop();
            }
        }
        else if constexpr(is_specialization<std::decay_t<T>, std::queue>::value)
        {
            T tmpValue = value;
            while(!tmpValue.empty())
            {
                m_value.push_back(tmpValue.front());
                tmpValue.pop();
            }
        }
        else
        {
            for(const auto& v : value)
            {
                m_value.push_back(v);
            }
        }
    }

public:
    Vector() = default;

    template<class T>
    requires VectorValue<T> // TODO
    Vector(T&& value) { initFromContainer(value); }

    Type type() const override;
    bool isEmpty() const override;
    std::size_t size() const override;

    template<class T>
    requires VectorValue<T> // TODO
    void operator=(T&& value) { initFromContainer(value); }

    // TODO: handle all cases
    template<class T>
    requires VectorValue<T>
    operator T()
    {
        T result;

        if constexpr(is_specialization<std::decay_t<T>, std::stack>::value || is_specialization<std::decay_t<T>, std::queue>::value)
        {
            for(auto& v : m_value)
                result.push(v);
        }
        else if constexpr(is_specialization<std::decay_t<T>, std::set>::value || is_specialization<std::decay_t<T>, std::multiset>::value)
        {
            std::transform(m_value.begin(), m_value.end(), std::inserter(result.begin(), result.end()),
                           [](auto& v){return v;});
        }
        else if constexpr(is_specialization<std::decay_t<T>, std::forward_list>::value)
        {
            result.assign(m_value.cbegin(), m_value.cend());
        }
        else
        {
            for(auto& v : m_value)
                result.push_back(v);
        }

        return result;
    }

    ValueType& front() { return m_value.front(); }
    const ValueType& front() const { return m_value.front(); }

    ValueType& back() { return m_value.back(); }
    const ValueType& back() const { return m_value.back(); }

    auto begin() { return Iterator<ValueType>(0, m_value.begin(), m_value.end()); }
    auto end() { return Iterator<ValueType>(m_value.size(), m_value.end(), m_value.end()); }

    template<class T>
    requires IntegerValue<T>
    ValueType& at(T&& index) { return m_value.at(index); }

    template<class T>
    requires IntegerValue<T>
    const ValueType& at(T&& index) const { return m_value.at(index); }

    void push_back(ValueType&&);
    void push_back(const ValueType&);

    void push_front(ValueType&&);
    void push_front(const ValueType&);
};

}

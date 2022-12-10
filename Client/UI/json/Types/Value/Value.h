#pragma once

#include <memory>
#include <ostream>

#include "../Base.hpp"
#include "../Bool/Bool.h"
#include "../Integer/Integer.h"
#include "../Map/Map.h"
#include "../Null/Null.h"
#include "../Number/Number.h"
#include "../String/String.h"
#include "../Vector/Vector.h"
#include "../Iterator/Iteartor.h"

#include <iostream>

namespace Types {

// TODO: add comparation between Base types
// TODO: add casting for Base types
class Value
{
    std::unique_ptr<Base> m_data;

    using VectorType = Vector<Value>;
    using MapType = Map<Value>;

    template<class T>
    void assign(T&& value)
    {
        if constexpr(IntegerValue<T>) m_data.reset(new Integer(value));
        else if constexpr(NumberValue<T>) m_data.reset(new Number(value));
        else if constexpr(BoolValue<T>) m_data.reset(new Bool(value));
        else if constexpr(StringValue<T>) m_data.reset(new String(value));
        else if constexpr(VectorValue<T>) m_data.reset(new VectorType(value));
        else if constexpr(NullValue<T>) m_data.reset(new Null(nullptr));
        else m_data.reset(new Null());
    }

public:
    explicit Value();

    Value(const std::initializer_list<Value>& l)
    {
        bool isObject = std::all_of(l.begin(), l.end(), [](const Value& v)
        {
            return v.isVector() && v.size() == 2 && v.front().isString();
        });

        if(isObject)
        {
            m_data.reset(new MapType());
            for(const auto& v : l)
            {
                std::string s = v.front();
                operator[](s) = v.back();
            }
        } else {
            m_data.reset(new VectorType(l));
        }
    }
    Value(const Value& other);
    Value& operator=(const Value& other);
    Value(Value&&) = default;
    Value& operator=(Value&&) = default;
    ~Value() = default;

    template<class T, class = std::enable_if_t<!std::is_same_v<std::decay_t<T>, Value>>>
    Value(T&& value) { assign(value); }

    template<class T, class = std::enable_if_t<!std::is_same_v<std::decay_t<T>, Value>>>
    void operator=(T&& value)
    {
       assign(value);
    }

    template<class T>
    requires StringValue<T>
    Value& operator[](T&& key)
    {
        if(isNull()) m_data.reset(new MapType());
        if(isMap()) return static_cast<MapType*>(m_data.get())->at(key);

        throw std::runtime_error("operator[](StringValue key)");
    }

    template<class T>
    requires IntegerValue<T>
    Value& operator[](T&& index)
    {
        if(isVector()) return static_cast<VectorType*>(m_data.get())->at(index);
    }

    template<class T>
    operator T() const
    {
        if constexpr(IntegerValue<T>) return static_cast<T>(*static_cast<Integer*>(m_data.get()));
        else if constexpr(NumberValue<T>) return static_cast<T>(*static_cast<Number*>(m_data.get()));
        else if constexpr(BoolValue<T>) return static_cast<T>(*static_cast<Bool*>(m_data.get()));
        else if constexpr(StringValue<T>) return static_cast<T>(*static_cast<String*>(m_data.get()));
        else if constexpr(VectorValue<T>) return static_cast<T>(*static_cast<VectorType*>(m_data.get()));
        else if constexpr(NullValue<T>) return static_cast<T>(*static_cast<Null*>(m_data.get()));
        else return T();
    }

    Value &front();
    const Value& front() const;

    Value& back();
    const Value& back() const;

    void push_back(Value&&);
    void push_back(const Value&);

    void push_front(Value&&);
    void push_front(const Value&);

    bool isInteger() const;
    bool isNumber() const;
    bool isBool() const;
    bool isString() const;
    bool isVector() const;
    bool isMap() const;
    bool isNull() const;

    bool isEmpty() const;
    std::size_t size() const;

    static Value object() { return Value({}); }
    static Value array(std::initializer_list<Value> l = {}) { return Value(std::vector<Value>(l)); }
    static Value parse(const std::string&);

    std::string dump();

    Iterator<Value> begin()
    {
        if(isMap()) return static_cast<MapType*>(m_data.get())->begin();
        if(isVector()) return static_cast<VectorType*>(m_data.get())->begin();

        std::cout << (int)m_data->type() << std::endl;

        throw std::runtime_error("begin error");
    }

    Iterator<Value> end()
    {
        if(isMap()) return static_cast<MapType*>(m_data.get())->end();
        if(isVector()) return static_cast<VectorType*>(m_data.get())->end();

        throw std::runtime_error("end error");
    }

    friend std::ostream& operator <<(std::ostream& os, Value& p)
    {
        os << p.dump();
        return os;
    }
};

}

using json = Types::Value;

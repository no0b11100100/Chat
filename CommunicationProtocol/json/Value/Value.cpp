#include "Value.h"
#include "Dumper.hpp"
#include "Parser.hpp"

using namespace Types;

Value::Value()
    : m_data{new Null()}
{}

Value::Value(const Value &other)
{
    *this = other;
}

Value &Value::operator=(const Value &other)
{
    if(other.isInteger())
        m_data.reset(new Integer(*(static_cast<Integer*>(other.m_data.get()))));
    else if(other.isNumber())
        m_data.reset(new Number(*(static_cast<Number*>(other.m_data.get()))));
    else if(other.isBool())
        m_data.reset(new Bool(*(static_cast<Bool*>(other.m_data.get()))));
    else if(other.isString())
        m_data.reset(new String(*(static_cast<String*>(other.m_data.get()))));
    else if(other.isVector())
        m_data.reset(new Vector<Value>(*(static_cast<Vector<Value>*>(other.m_data.get()))));
    else if(other.isMap())
        m_data.reset(new Map<Value>(*(static_cast<Map<Value>*>(other.m_data.get()))));
    else
        m_data.reset(new Null(*(static_cast<Null*>(other.m_data.get()))));

    return *this;
}

Value &Value::front()
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->front();
    throw std::runtime_error("front error");
}

const Value &Value::front() const
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->front();

    throw std::runtime_error("front error");
}

Value &Value::back()
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->back();

    throw std::runtime_error("back error");
}

const Value &Value::back() const
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->back();
    throw std::runtime_error("back error");
}

void Value::push_back(Value&& value)
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->push_back(value);
}

void Value::push_back(const Value& value)
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->push_back(value);
}

void Value::push_front(Value&& value)
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->push_front(value);
}

void Value::push_front(const Value& value)
{
    if(isVector()) return static_cast<Vector<Value>*>(m_data.get())->push_front(value);
}

bool Value::isInteger() const
{
    return m_data->type() == Type::INTEGER;
}

bool Value::isNumber() const
{
    return m_data->type() == Type::NUMBER;
}

bool Value::isBool() const
{
    return m_data->type() == Type::BOOL;
}

bool Value::isString() const
{
    return m_data->type() == Type::STRING;
}

bool Value::isVector() const
{
    return m_data->type() == Type::ARRAY;
}

bool Value::isMap() const
{
    return m_data->type() == Type::MAP;
}

bool Value::isNull() const
{
    return m_data->type() == Type::NIL;
}

bool Value::isEmpty() const
{
    return m_data->isEmpty();
}

std::size_t Value::size() const
{
    return m_data->size();
}

Value Value::parse(const std::string & json)
{
    Parser parser;
    return parser.parse(json);
}

std::string Value::dump()
{
    Dumper dumper;
    return dumper.dump(*this);
}

std::string Value::prettyDump()
{
    Dumper dumper;
    return dumper.prettyDump(*this);
}
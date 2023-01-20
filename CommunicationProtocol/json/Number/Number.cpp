#include "Number.h"

Types::Number::Number()
    : m_value{0.0}
{}

Types::Type Types::Number::type() const
{
    return Type::NUMBER;
}

bool Types::Number::isEmpty() const
{
    // TODO: add exception
    return true;
}

std::size_t Types::Number::size() const
{
    // TODO: add exception
    return 0;
}

std::ostream& Types::operator <<(std::ostream & os, const Number& i)
{
    os << std::to_string(i.m_value);
    return os;
}

std::istream& Types::operator >>(std::istream & is, Number& value)
{
    is >> value.m_value;
    return is;
}

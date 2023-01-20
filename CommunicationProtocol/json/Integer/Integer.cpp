#include "Integer.h"

Types::Integer::Integer()
    : m_value{0}
{}

Types::Type Types::Integer::type() const
{
    return Type::INTEGER;
}


bool Types::Integer::isEmpty() const
{
    // TODO: add exception
    return true;
}

std::size_t Types::Integer::size() const
{
    // TODO: add exception
    return 0;
}

std::ostream& Types::operator <<(std::ostream & os, const Integer& i)
{
    os << std::to_string(i.m_value);
    return os;
}

std::istream& Types::operator >>(std::istream & is, Integer& value)
{
    is >> value.m_value;
    return is;
}

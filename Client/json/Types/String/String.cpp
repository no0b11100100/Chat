#include "String.h"

Types::String::String()
    : m_value{""}
{}

Types::Type Types::String::type() const
{
    return Type::STRING;
}

bool Types::String::isEmpty() const
{
    return m_value.empty();
}

std::size_t Types::String::size() const
{
    return m_value.size();
}

Types::String Types::String::substr(size_t pos, size_t len) const
{
    return m_value.substr(pos, len);
}

std::ostream& Types::operator <<(std::ostream & os, const String& i)
{
    os << i.m_value;
    return os;
}

std::istream& Types::operator >>(std::istream & is, String& value)
{
    is >> value.m_value;
    return is;
}

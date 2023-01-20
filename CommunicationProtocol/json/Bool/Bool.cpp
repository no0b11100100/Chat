#include "Bool.h"

Types::Bool::Bool()
    : m_value{false}
{}

Types::Type Types::Bool::type() const
{
    return Type::BOOL;
}

bool Types::Bool::isEmpty() const
{
    // TODO: add exception
    return true;
}

std::size_t Types::Bool::size() const
{
    // TODO: add exception
    return 0;
}

std::ostream& Types::operator <<(std::ostream & os, const Bool& i)
{
    os << std::boolalpha << std::to_string(i.m_value);
    return os;
}

std::istream& Types::operator >>(std::istream & is, Bool& value)
{
    is >> value.m_value;
    return is;
}

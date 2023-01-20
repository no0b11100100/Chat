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

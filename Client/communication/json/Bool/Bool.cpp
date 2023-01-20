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

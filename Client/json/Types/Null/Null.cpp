#include "Null.h"

Types::Type Types::Null::type() const
{
    return Type::NIL;
}

bool Types::Null::isEmpty() const
{
    // TODO: add exception
    return true;
}

std::size_t Types::Null::size() const
{
    // TODO: add exception
    return 0;
}

std::ostream& Types::operator <<(std::ostream & os, const Null&)
{
    os << "null";
    return os;
}

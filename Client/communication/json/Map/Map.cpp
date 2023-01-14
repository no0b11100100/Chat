#include "Map.h"
#include "../Value/Value.h"

using namespace Types;

template<class T>
Type Map<T>::type() const
{
    return Type::MAP;
}

template <class T>
bool Map<T>::isEmpty() const
{
    return m_value.empty();
}

template<class T>
std::size_t Map<T>::size() const
{
    return m_value.size();
}

template<class T>
bool Map<T>::contains(const std::string &key)
{
    return find(key) != m_value.end();
}

template class Types::Map<Value>;

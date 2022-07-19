#include "Vector.h"
#include "../Value/Value.h"

using namespace Types;

template<class T>
Type Vector<T>::type() const
{
    return Type::ARRAY;
}

template<class T>
bool Vector<T>::isEmpty() const
{
    return m_value.empty();
}

template<class T>
std::size_t Vector<T>::size() const
{
    return m_value.size();
}

template<class T>
void Vector<T>::push_back(T&& value)
{
    m_value.push_back(value);
}

template<class T>
void Vector<T>::push_back(const T& value)
{
    m_value.push_back(value);
}

template<class T>
void Vector<T>::push_front(T&& value)
{
    m_value.push_front(value);
}

template<class T>
void Vector<T>::push_front(const T& value)
{
    m_value.push_front(value);
}

//template<class T>
//void Vector<T>::unique()
//{
//    auto last = std::unique(m_value.begin(), m_value.end());
//    m_value.erase(last, m_value.end());
//}

template class Types::Vector<Value>;

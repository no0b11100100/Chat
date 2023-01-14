#pragma once
#include <iostream>
#include <type_traits>
#include <vector>
#include <list>
#include <forward_list>
#include <set>
#include <deque>
#include <stack>
#include <queue>
#include <initializer_list>

namespace Types {

enum class Type
{
    INTEGER,
    NUMBER,
    BOOL,
    STRING,
    CHAR,
    ARRAY,
    MAP,
    NIL,
    FUNCTION,
};

class Base
{
public:
    virtual Type type() const = 0;
    virtual bool isEmpty() const = 0;
    virtual std::size_t size() const = 0;

    virtual ~Base() = default;
};

template<class T>
concept BoolValue =
        std::is_same_v<std::decay_t<T>, bool>;

template<class T>
concept IntegerValue =
        std::is_same_v<std::decay_t<T>, short> ||
        std::is_same_v<std::decay_t<T>, unsigned short> ||
        std::is_same_v<std::decay_t<T>, long> ||
        std::is_same_v<std::decay_t<T>, int> ||
        std::is_same_v<std::decay_t<T>, long long> ||
        std::is_same_v<std::decay_t<T>, unsigned> ||
        std::is_same_v<std::decay_t<T>, unsigned long> ||
        std::is_same_v<std::decay_t<T>, unsigned long long> ||
        std::is_same_v<std::decay_t<T>, std::size_t>;

template<class T>
concept NullValue =
        std::is_same_v<std::decay_t<T>, std::nullptr_t>;

template<class T>
concept NumberValue =
        std::is_same_v<std::decay_t<T>, float> ||
        std::is_same_v<std::decay_t<T>, double> ||
        std::is_same_v<std::decay_t<T>, long double>;

template<class T>
concept CharValue =
        std::is_same_v<std::decay_t<T>, char>;

template<class T>
concept StringValue =
        std::is_same_v<std::decay_t<T>, std::string> ||
        std::is_same_v<std::decay_t<T>, std::string_view> ||
        std::is_same_v<std::decay_t<T>, std::wstring> ||
        std::is_same_v<std::decay_t<T>, std::wstring_view> ||
        std::is_same_v<std::decay_t<T>, const char*>;

template<class T>
concept ScalarValue =
        IntegerValue<T> &&
        BoolValue<T> &&
        StringValue<T> &&
        CharValue<T> &&
        NumberValue<T> &&
        NullValue<T>;

//template<class T>
//concept VoidValue =
//        std::is_same_v<std::decay_t<T>, void>;

template <class T, class R = void>
concept FunctionValue = std::is_invocable_v<T, R>;

template<typename Test, template<typename...> class Ref>
struct is_specialization : std::false_type {};

template<template<typename...> class Ref, typename... Args>
struct is_specialization<Ref<Args...>, Ref>: std::true_type {};

template<class T>
concept VectorValue =
        is_specialization<std::decay_t<T>, std::vector>::value ||
        is_specialization<std::decay_t<T>, std::list>::value ||
        is_specialization<std::decay_t<T>, std::forward_list>::value ||
        is_specialization<std::decay_t<T>, std::set>::value ||
        is_specialization<std::decay_t<T>, std::multiset>::value ||
        is_specialization<std::decay_t<T>, std::deque>::value ||
        is_specialization<std::decay_t<T>, std::stack>::value ||
        is_specialization<std::decay_t<T>, std::initializer_list>::value ||
        is_specialization<std::decay_t<T>, std::queue>::value;
}

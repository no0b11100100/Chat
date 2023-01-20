cmake_minimum_required(VERSION 3.5)

project(ValueSystem LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(ValueSystem
    main.cpp

    Base.hpp

    Integer/Integer.h
    Integer/Integer.cpp

    Number/Number.h
    Number/Number.cpp

    Bool/Bool.h
    Bool/Bool.cpp

    String/Char.hpp
    String/String.h
    String/String.cpp

    Null/Null.h
    Null/Null.cpp

    Value/Value.h
    Value/Value.cpp

    Vector/Vector.h
    Vector/Vector.cpp

    Map/Map.h
    Map/Map.cpp
    Map/Pair.hpp

    Iterator/Iteartor.h

    Value/Dumper.hpp

    Value/Separators.hpp
    Value/Parser.hpp
)

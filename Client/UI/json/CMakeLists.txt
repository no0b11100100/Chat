cmake_minimum_required(VERSION 3.5)

project(ValueSystem LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(ValueSystem
    main.cpp

    Types/Base.hpp

    Types/Integer/Integer.h
    Types/Integer/Integer.cpp

    Types/Number/Number.h
    Types/Number/Number.cpp

    Types/Bool/Bool.h
    Types/Bool/Bool.cpp

    Types/String/String.h
    Types/String/String.cpp

    Types/Null/Null.h
    Types/Null/Null.cpp

    Types/Value/Value.h
    Types/Value/Value.cpp

    Types/Vector/Vector.h
    Types/Vector/Vector.cpp

    Types/Map/Map.h
    Types/Map/Map.cpp
    Types/Map/Pair.hpp

    Types/Iterator/Iteartor.h

    Types/Value/Dumper.hpp

    Types/Value/Separators.hpp
    Types/Value/Parser.hpp
)

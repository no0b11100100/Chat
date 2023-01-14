#pragma once
#include <functional>
#include <vector>
#include <sstream>

#include "Value.h"
#include "../Bool/Bool.h"
#include "../Integer/Integer.h"
#include "../Map/Map.h"
#include "../Null/Null.h"
#include "../Number/Number.h"
#include "../String/String.h"
#include "../Vector/Vector.h"
#include "Separators.hpp"

class Parser
{
    enum class TokenValue { Key, Value };
    enum class SlashStatus { Start, InProgress, End, Invalid };

    SlashStatus increment_if(bool condition, const SlashStatus& s)
    {
        if(!condition) return s;

        switch(s)
        {
        case SlashStatus::Start: return SlashStatus::InProgress;
        case SlashStatus::InProgress: return SlashStatus::End;
        case SlashStatus::End: return SlashStatus::Invalid;
        default: return SlashStatus::Invalid;
        }
    }

    Types::Value json;
    std::vector<Types::Value*> tree{&json};
    using Pair = std::pair<char, std::function<void()>>;
    std::vector<Pair> m_handlers;
    std::vector<Types::Value> array;
    Separators m_separators;

    std::string key;
    std::string value;
    TokenValue tokenValue = TokenValue::Key;
    short quotesCount{0};
    bool isStringValueReading{false};
    char prevSeparator = 0;
//    bool isRoot{false};
    SlashStatus slashStatus{SlashStatus::Start};

    Types::Value setCorrectType(const std::string& value)
    {
        static std::array<char, 10> numbers{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
        if(value == "true")
            return Types::Value(true);
        if(value == "false")
            return Types::Value(false);
        if(value == "null")
            return Types::Value(nullptr);
        if(value.empty())
            return Types::Value();

        if(std::all_of(value.front() == '-' ? std::next(value.cbegin()) : value.cbegin(), value.cend(), [&](const char& c){
                       return std::any_of(numbers.cbegin(), numbers.cend(), [c](const char& num){ return c == num; });
        }))
        {
            std::istringstream ss{value};
            if(value.front() == '-')
            {
                long long v;
                ss >> v;
                return Types::Value(v);
            }

            size_t v;
            ss >> v;
            return Types::Value(v);
        }
        if(std::all_of(value.front() == '-' ? std::next(value.cbegin()) : value.cbegin(), value.cend(), [&](const char& c){
            static short dotCounter{0};
            if(c == '.')
            {
                if(dotCounter > 1) return false;
                ++dotCounter;
                return true;
            }
            return std::any_of(numbers.cbegin(), numbers.cend(), [c](const char& num){ return c == num; });
        }))
        {
            return Types::Value(std::stold(value));
        }

        return Types::Value(value);
    }

    void setPair()
    {
        if(!key.empty())
            tree.back()->operator[](key) = setCorrectType(value);
        value = "";
        key = "";
    }

    // {
    void handleBraceIn()
    {
        if(tree.size() == 1 && tree.back()->isNull())
        {
            tokenValue = TokenValue::Key;
            *tree.back() = Types::Value::object();
        }
        if(tree.back()->isVector())
        {
            std::cout << "vector\n";
            tokenValue = TokenValue::Key;
            tree.back()->push_back(Types::Value::object());
            tree.push_back(&tree.back()->back()); //(&static_cast<Types::Vector<Types::Value>*>(tree.back()->m_data.get())->back());
        }
        else if(!key.empty())
        {
            tokenValue = TokenValue::Key;
            tree.back()->operator[](key) = Types::Value::object();
            tree.push_back(&tree.back()->operator[](key));
            key = "";
        }
    }

    // }
    void handleBraceOut()
    {
        setPair();
        tree.erase(std::prev(tree.end()));
        quotesCount = 0;
        key = "";
    }

    // [
    void handleSquareBracketIn()
    {
        std::array<char, 2> allowedSeparators{',', ':'};
        if((std::none_of(allowedSeparators.cbegin(), allowedSeparators.cend(), [s = prevSeparator](char c){ return c == s; }))) return;
//            throw ParseError("error after [");

        tokenValue = TokenValue::Key;
        tree.back()->operator[](key) = Types::Value::array();
        tree.push_back(&tree.back()->operator[](key));
        key = "";
    }

    // ]
    void handleSquareBracketOut()
    {
        if(!key.empty()) tree.back()->push_back(setCorrectType(key));
        quotesCount = 0;
        tree.erase(std::prev(tree.end()));
        key = "";
    }

    // :
    void handleColon()
    {
        if(prevSeparator != '"') return;//throw ParseError("error after :");
        tokenValue = TokenValue::Value;
        quotesCount = 0;
    }

    // "
    void handleQuotationMark()
    {
        ++quotesCount;
        if(quotesCount > 2) std::cout << "bad\n";
        isStringValueReading = tokenValue == TokenValue::Value && quotesCount == 1;
    }

    // ,
    void handleComma()
    {
        if(tree.back()->isVector())
        {
            if(!key.empty()) tree.back()->push_back(setCorrectType(key));
        }
        else setPair();
        tokenValue = TokenValue::Key;
        quotesCount = 0;
        key = "";
    }

    void readToken(char symbol)
    {
        slashStatus = increment_if(symbol == '\\', slashStatus);
        bool isEndReadingString = value.back() != '\\' && symbol == '"';
        if(isStringValueReading && isEndReadingString) handleQuotationMark();
        else if(tokenValue == TokenValue::Key) key += symbol;
        else if(tokenValue == TokenValue::Value)
        {
            bool isEscapedNotSlash = (slashStatus == SlashStatus::InProgress && symbol != '\\');
            bool isEndHandlingEscapedSlash = (slashStatus == SlashStatus::End || isEscapedNotSlash);
            if(slashStatus == SlashStatus::InProgress && m_separators.find({'\\', symbol}))
            {
                value += symbol;
                slashStatus = SlashStatus::Start;
            }
            else if(isEndHandlingEscapedSlash && m_separators.find(symbol))
            {
                slashStatus = SlashStatus::Start;
                value.back() = symbol;
            }
            else value += symbol;
        }
    }

    void handleWhiteSpace()
    {
        if (prevSeparator == '"' && quotesCount == 1)
        {
            key += " ";
        }
    }

public:
    Parser()
    {
        m_handlers = {
            {'{', [&](){ handleBraceIn(); }},
            {'}', [&](){ handleBraceOut(); }},
            {'[', [&](){ handleSquareBracketIn(); }},
            {']', [&](){ handleSquareBracketOut(); }},
            {':', [&](){ handleColon(); }},
            {'"', [&](){ handleQuotationMark(); }},
            {',', [&](){ handleComma(); }},
            {'\n', [&](){}},
            {'\t', [&](){}},
            {'\\', [&](){}},
            {' ', [&](){ handleWhiteSpace(); }}
        };
    }

    Types::Value parse(const std::string& str)
    {
        if(str.front() == '{' || str.front() == '[') {
            for(const auto& symbol : str)
            {
                if(auto it = std::find_if(m_handlers.cbegin(), m_handlers.cend(), [&symbol](const Pair& p){ return p.first == symbol; });
                        it != m_handlers.cend() && !isStringValueReading)
                {
                    it->second();
                    if (it->first != ' ') prevSeparator = it->first;
                }
                else
                    readToken(symbol);
            }
        } else {
            json = setCorrectType(str);
            tree.erase(tree.begin());
        }

        return json;
    }
};

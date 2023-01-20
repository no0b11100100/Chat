#pragma once
#include <string>
#include <vector>
#include <array>
#include <cassert>

class Separators
{
    using ValueType = std::pair<char, std::string>;
    std::vector<ValueType> m_separators{
        {'\n', "\\n"},
        {'\t', "\\t"},
        {'\"', "\\\""},
        {'\\', "\\\\"}
    };

    char escapeSymbols(const std::array<char, 2>& s)
    {
        assert(!s.empty() && s[0] == '\\');
        switch(s[1])
        {
        case 't': return '\t';
        case 'n': return '\n';
        default: return ' ';
        }
    }
public:
    bool find(char c)
    {
        return std::find_if(m_separators. cbegin(), m_separators.cend(), [&c](const ValueType& p){
            return p.first == c;
        }) != m_separators.cend();
    }
    bool find(std::array<char, 2> s) { return find(escapeSymbols(s)); }
    std::string at(char c)
    {
        return std::find_if(m_separators. cbegin(), m_separators.cend(), [&c](const ValueType& p){
            return p.first == c;
        })->second;
    }
    std::string at(std::array<char, 2> s) { return at(escapeSymbols(s)); }
};

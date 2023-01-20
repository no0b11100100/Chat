#pragma once
#include <string>
#include <cassert>

#include "Value.h"
#include "Separators.hpp"

class Dumper
{
    std::string m_data;
    std::size_t m_indents;

    void addIndents()
    {
        for(std::size_t i{0}; i < m_indents; ++i)
            m_data += "\t";
    }

    void normilize()
    {
        auto index = m_data.find_last_of(",");
        if(index < m_data.size())
            m_data = m_data.substr(0, index);
    }

    void normilizeEmpty()
    {
        if(!m_data.empty() && m_data.front() == '{' && std::all_of(std::next(m_data.cbegin()), std::prev(m_data.cend()), [](const char c)
        { return c == '\n' || c == '\t'; }))
        {
            m_data = "{}";
        }

        if(!m_data.empty() && m_data.front() == '[' && std::all_of(std::next(m_data.cbegin()), std::prev(m_data.cend()), [](const char c)
        { return c == '\n' || c == '\t'; }))
        {
            m_data = "[]";
        }
    }

public:
    Dumper()
        : m_data{""},
          m_indents{0}
    {}

    std::string prettyDump(Types::Value& data)
    {
        if(data.isNull())
        {
            m_data += "null";
        }
        else if (data.isBool())
        {
            bool b = data;
            m_data += std::to_string(b);
        }
        else if (data.isInteger())
        {
            long long i = data;
            m_data += std::to_string(i);
        }
        else if (data.isNumber())
        {
            long double d = data;
            m_data += std::to_string(d);
        }
        else if(data.isString())
        {
            std::string s = data;
            m_data += "\"";
            Separators separators;
            for(auto it = s.cbegin(); it != s.cend(); ++it)
            {
                if(separators.find(*it))
                {
                    if(*it == '\\' &&  separators.find({*it, *std::next(it)}))
                    {
                        m_data += separators.at({*it, *std::next(it)});
                        it = std::next(it);
                    }
                    else m_data += separators.at(*it);
                }
                else m_data += *it;
            }

            m_data += "\"";
        }
        else if (data.isVector())
        {
            m_data += "[\n";
            ++m_indents;
            for(auto it = data.begin(); it != data.end(); ++it)
            {
                addIndents();
                prettyDump(it.value());
                m_data += ",\n";
            }
            normilize();
            m_data += "\n";
            --m_indents;
            addIndents();
            m_data += "]";
            normilizeEmpty();
        }
        else if(data.isMap())
        {
            m_data += "{\n";
            ++m_indents;
            for(auto it = data.begin(); it != data.end(); ++it)
            {
                std::string key = it.key();
                addIndents();
                m_data += "\"" + key + "\":";
                prettyDump(it.value());
                m_data += ",\n";
            }
            normilize();
            m_data += "\n";
            --m_indents;
            addIndents();
            m_data += "}";
            normilizeEmpty();
        }
        else
        {
            std::cout << "Unknown type";
            return m_data;
        }
        return m_data;
    }

    std::string dump(Types::Value& data)
    {
        if(data.isNull())
        {
            m_data += "null";
        }
        else if (data.isBool())
        {
            bool b = data;
            m_data += std::to_string(b);
        }
        else if (data.isInteger())
        {
            long long i = data;
            m_data += std::to_string(i);
        }
        else if (data.isNumber())
        {
            long double d = data;
            m_data += std::to_string(d);
        }
        else if(data.isString())
        {
            std::string s = data;
            m_data += "\"";
            Separators separators;
            for(auto it = s.cbegin(); it != s.cend(); ++it)
            {
                if(separators.find(*it))
                {
                    if(*it == '\\' &&  separators.find({*it, *std::next(it)}))
                    {
                        m_data += separators.at({*it, *std::next(it)});
                        it = std::next(it);
                    }
                    else m_data += separators.at(*it);
                }
                else m_data += *it;
            }

            m_data += "\"";
        }
        else if (data.isVector())
        {
            m_data += "[";
            for(auto it = data.begin(); it != data.end(); ++it)
            {
                dump(it.value());
                m_data += ",";
            }
            normilize();
            m_data += "]";
        }
        else if(data.isMap())
        {
            m_data += "{";
            ++m_indents;
            for(auto it = data.begin(); it != data.end(); ++it)
            {
                std::string key = it.key();
                m_data += "\"" + key + "\":";
                dump(it.value());
                m_data += ",";
            }
            normilize();
            m_data += "}";
        }
        else
        {
            std::cout << "Unknown type";
            return m_data;
        }
        return m_data;
    }
};

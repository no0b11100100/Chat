#pragma once

#include <iostream>
#include <string_view>
#include <source_location>
#include <ctime>

namespace Logger
{
    namespace {
        using LOG_LEVEL = std::string_view;
        constexpr LOG_LEVEL INFO = "[INFO]:";
        constexpr LOG_LEVEL WARNING = "[WARNING]:";
        constexpr LOG_LEVEL ERROR = "[ERROR]:";

        using COLOR = std::string_view;
        constexpr COLOR INFO_COLOR  = "\033[0m";
        constexpr COLOR ERROR_COLOR    = "\033[31m";
        constexpr COLOR WARNING_COLOR = "\033[33m";

        std::string parseFileName(const std::string& fileName)
        {
            auto it = std::find(fileName.crbegin(), fileName.crend(), '/');
            std::string result;
            std::copy(fileName.crbegin(), it, std::back_inserter(result));
            std::reverse(result.begin(), result.end());
            return result;
        }

        std::string currentDateTime() {
            std::time_t t = std::time(nullptr);
            std::tm* now = std::localtime(&t);

            char buffer[128];
            strftime(buffer, sizeof(buffer), "%X", now);
            return buffer;
        }

        class Printer {
            private:
            std::string call_info;

            template <typename... Args>
            void print(const COLOR color, const LOG_LEVEL log_level, Args&&... args)
            {
                std::cout << color << currentDateTime() << " " << call_info << ": " << log_level << std::boolalpha;
                ((std::cout << ' ' << std::forward<Args>(args)), ...);
                std::cout << INFO_COLOR << std::endl;
            }

            public:

            Printer(const std::source_location location)
            {
                call_info = parseFileName(location.file_name()) + ":" + std::to_string(location.line());
            }

            template <typename... Args>
            void Info(Args&&... args)
            {
                print(INFO_COLOR, INFO, args...);
            }

            template <typename... Args>
            void Warning(Args&&... args)
            {
                print(WARNING_COLOR, WARNING, args...);
            }

            template <typename... Args>
            void Error(Args&&... args)
            {
                print(ERROR_COLOR, ERROR, args...);
            }

        };
    }

    auto log(const std::source_location location = std::source_location::current())
    {
        Printer printer(location);
        return printer;
    };

};

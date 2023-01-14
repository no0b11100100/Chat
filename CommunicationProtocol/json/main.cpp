#include <iostream>

#include "Value/Value.h"
using namespace Types;
using namespace std;

// #include "String/Char.hpp"
// #include "Value/Value.h"

// class Data : public ClassParser
// {
//     int val;
//     std::string s;
// public:

//     Data()
//     : val{10},
//     s{"some string"}
//     {}

//     void reset() {
//         val = 20;
//         s = "reset";
//     }

//     void print()
//     {
//         std::cout << val << " " << s << std::endl;
//     }

//     virtual Value toJson() override
//     {
//         Value js({});
//         js["value"] = val;
//         js["data"] = s;
//         return js;
//     }

//     virtual void fromJson(Value js) override
//     {
//         val = js["value"];
//         s = js["data"];
//     }
// };


int main()
{
//    Value js;
//    js["data"] = 10;
//    js["value"] = 423.12;
//    js["array"] = Value::array({10,12,4, Value::object()});
//    js["array"][3]["value"] = "string value";
//    js["array"][3]["integer"] = 13123;
//    js["stub"] = nullptr;
//    js["string"] = "string value";

//    auto it = js.begin();
//    int i = it.value();
//    std::string s = it.key();
//    cout << s << " " << it.value().isInteger() << endl;
//    cout << i << endl;

//    ++it;
//    std::string sss = it.key();
//    double d = it.value();
//    cout << sss << " " << d << endl;

//    *it = 12.4;
//    d = js["value"];
//    cout << sss << " " << d << endl;

//    ++it;
//    std::string ss = it.key();
//    cout << ss << " " << it.value().isVector() << endl;
//    std::vector<int> v = it.value();
//    for(auto val : v) { cout << val << endl; }

//    cout << js.dump() << endl;

//    cout << "pass\n";

    // Value js({});
    // cout << js << endl;







    // Data d;
    // d.print();
    // Value js;

    // js = d;

    // cout << js << endl;

    // d.reset();
    // d.print();

    // d = js;

    // d.print();

    // const char * v = js["data"];

    // Value js;
    // js = Value::array({1,2,3});

    // std::vector<int> v;
    // v = static_cast<vector<int>>(js);

    Value js;
    js = Value::array({"s", "a", "m"});

    std::string s = js[0];

    cout << s;

    // std::string str = "value";
    // String<Types::Value> s(str);
    // s.print();


    return 0;
}

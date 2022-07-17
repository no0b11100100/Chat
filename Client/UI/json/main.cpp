#include <iostream>

#include "Types/Value/Value.h"
using namespace Types;
using namespace std;
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

    Value js({});
    cout << js << endl;

    return 0;
}

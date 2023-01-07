#include <iostream>

#include "Value/Value.h"
using namespace Types;
using namespace std;

class Data : public ClassParser
{
    int val;
public:

    Data()
    : val{10}
    {}

    void reset() {
        val = 20;
    }

    void print()
    {
        std::cout << val << std::endl;
    }

    virtual Value toJson() override
    {
        Value js({});
        js["value"] = val;
        return js;
    }

    virtual void fromJson(Value js) override
    {
        val = js["value"];
    }
};


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

    Data d;
    d.print();
    Value js;
    js = d;

    cout << js << endl;

    d.reset();
    d.print();

    d = js;

    d.print();




    return 0;
}

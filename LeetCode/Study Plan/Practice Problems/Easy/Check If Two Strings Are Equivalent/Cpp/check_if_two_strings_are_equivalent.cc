#include <iostream>
#include <vector>

typedef std::vector<std::string> vs;

bool check_if_two_strings_are_equivalent(vs& v1, vs& v2) {
    std::string s1 = "", s2 = "";
    for (int i=0;i<v1.size(); i++) s1 +=  v1[i];
    for (int i=0;i<v2.size(); i++) s2 +=  v2[i];
    return s1.compare(s2)==0;
}

int main () {
    vs v1{"ab", "c"}, v2{"a", "bc"};
    std::cout<<check_if_two_strings_are_equivalent(v1, v2)<<"\n";
}

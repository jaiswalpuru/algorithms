#include <iostream>
#include <map>

typedef std::map<char, int> mc;
typedef std::map<int, int> mi;

bool determine_if_two_strings_are_close(std::string s1, std::string s2) {
    if (s1.length() != s2.length()) return false;
    mc m1, m2;
    mi v1, v2;
    for (int i=0;i<s1.length(); i++) {
        m1[s1[i]]++;
        m2[s2[i]]++;
    }
    for (auto it : m1) {
        if (m2[it.first] == 0) return false;   
    }
    for (auto it : m1) v1[it.second]++;
    for (auto it : m2) v2[it.second]++;
    for (auto it : v1) {
        if (v2[it.first] != it.second) return false;
    }
    return true;
}

int main() {
    std::cout<<determine_if_two_strings_are_close("abc", "bca")<<"\n";
}

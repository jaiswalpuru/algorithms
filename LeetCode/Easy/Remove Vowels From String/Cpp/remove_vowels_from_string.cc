#include <iostream>

std::string remove_vowels_from_string(std::string s) {
    std::string str = "";
    for (int i=0; i<s.size(); i++) 
        if (!(s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u')) str += s[i];
    return str;
}

int main() {
    std::cout<<remove_vowels_from_string("leetcodeisacommunityforcoders")<<"\n";
}

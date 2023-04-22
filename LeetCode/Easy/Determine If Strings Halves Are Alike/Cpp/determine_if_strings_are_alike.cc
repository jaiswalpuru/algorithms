#include <iostream>
#include <cctype>

bool determine_if_strings_are_alike(std::string s) {
    int cnt = 0;
    transform(s.begin(), s.end(), s.begin(), ::toupper);
    for(int i=0; i<s.length()/2; i++)
        if (s[i] == 'A' || s[i] == 'E' || s[i] == 'O' || s[i] == 'I' || s[i] == 'U') cnt++;
    for(int i=s.length()/2; i<s.length(); i++) 
        if (s[i] == 'A' || s[i] == 'E' || s[i] == 'O' || s[i] == 'I' || s[i] == 'U') cnt--;
    return cnt == 0;
}

int main() {
    std::cout<<determine_if_strings_are_alike("book")<<"\n";
}

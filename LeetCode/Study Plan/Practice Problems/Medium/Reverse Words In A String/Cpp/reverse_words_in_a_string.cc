#include <iostream>
#include <vector>
#include <sstream>

std::string reverse_words_in_a_string(std::string s) {
    std::stringstream ss(s);
    std::string word;
	std::string res = "";
    while(ss>>word) {
        res = word + " " + res;
    }
    return res.substr(0, res.size()-1);
}

int main() {
    std::cout<<reverse_words_in_a_string("the sky is blue")<<"\n";
}

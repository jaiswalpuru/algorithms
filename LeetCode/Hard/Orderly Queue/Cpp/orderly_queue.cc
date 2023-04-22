#include <iostream>

std::string orderly_queue(std::string s, int k) {
    if (k == 1) {
        std::string ans = s;
        for (int i=0; i<s.length(); i++) {
            std::string temp = s.substr(i, s.length()) + s.substr(0, i);
            if (ans.compare(temp) > 0) {
                ans = temp;
            }
        }
        return ans;
    }else {
        sort(s.begin(), s.end());
        return s;
    }
}

int main() {
    std::string s = "cba";
    std::cout<<orderly_queue(s, 1)<<"\n";
}

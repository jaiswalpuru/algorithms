#include <iostream>
#include <unordered_map>
#include <climits>

void repetitions(std::string str) {
    std::unordered_map<char, int> char_cnt;
    int res = 1;
    int temp = 1;
    for (int i=1;i<str.length();i++){
        if (str[i] == str[i-1]) {
            temp++;
        }else {
            res = std::max(temp, res);
            temp =1;
        }
    }
    res = std::max(res, temp);
    std::cout<<res<<"\n";
}

int main() {
    std::string str;
    getline(std::cin, str);
    repetitions(str);
    return 0;
}

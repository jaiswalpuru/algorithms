#include <iostream>
#include <map>
#include <vector>

typedef std::vector<char> vc;
typedef std::map<int, vc> mic;
typedef std::map<char, int> mci;

std::string sort_characters_by_frequency(std::string s) {
    mci cnt;
    mic char_cnt;
    std::string res = "";
    for (int i=0; i<s.length(); i++) cnt[s[i]]++;
    for (auto it : cnt) char_cnt[it.second].push_back(it.first);
    for (auto it : char_cnt) {
        vc chars = it.second;
        for (int i=0; i<chars.size(); i++) {
            for (int j=0; j<it.first; j++) {
                res += chars[i];
            }
        }
    }
    reverse(res.begin(), res.end());
    return res;
}

int main() {
    std::cout<<sort_characters_by_frequency("tree")<<"\n";
}

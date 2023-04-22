#include <iostream>
#include <vector>
#include <map>

typedef std::string str;
typedef std::vector<str> vs;
typedef std::vector<vs> vss;
typedef std::map<str, vs> mss;

vss group_anagrams(vs v) {
    mss mp;
    for (int i=0;i<v.size(); i++) {
        str word = v[i];
        std::sort(word.begin(), word.end());
        mp[word].push_back(v[i]);
    }
    vss res;
    for (auto itr = mp.begin(); itr!= mp.end(); itr++) {
        res.push_back(itr->second);
    }
    return res;
}

int main() {
    vs v{"eat","tea","tan","ate","nat","bat"};
    vss res;
    res = group_anagrams(v);
    for (int i=0;i<res.size(); i++){
        for (int j=0;j<res[i].size();j++) std::cout<<res[i][j]<<" ";
        std::cout<<"\n";
    }
}

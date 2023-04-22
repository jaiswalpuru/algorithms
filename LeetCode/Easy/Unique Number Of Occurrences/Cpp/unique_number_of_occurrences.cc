#include <iostream>
#include <map>
#include <vector>

typedef std::vector<int> vi;
typedef std::map<int, int> mi;

bool unique_number_of_occurrences(vi v) {
    mi mp, cnt;
    for (int i=0; i<v.size(); i++) mp[v[i]]++;
    for (auto it=mp.begin(); it!=mp.end(); it++) {
        if (cnt[it->second] > 0) return false;
        cnt[it->second] = it->second;
    }
    return true;
}

int main() {
    std::cout<<unique_number_of_occurrences(vi{-3,0,1,-3,1,1,1,-3,10,0})<<"\n";
}

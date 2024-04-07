#include <iostream>
#include <queue>
#include <map>
#include <vector>
#include <string>
#include <utility>

typedef std::string st;
typedef std::vector<st> vs;
typedef std::queue<std::pair<st, vs>> qs;
typedef std::map<st, bool> msb;

vs stepword_chain(st start, st end, vs &dic) {
    msb has;
    vs ans;
    qs q;

    q.push({start, {start}});

    for (int i = 0; i < dic.size(); i++) has[dic[i]] = true;

    while(!q.empty()) {
        auto curr = q.front();
        q.pop();
        
        if (curr.first.compare(end) == 0) return curr.second;

        st s = "";
        for (int k = 0; k < curr.first.size(); k++) {
            for (int i = 0; i<26; i++) {
                s = curr.first.substr(0, k) + st(1, 'a' + i) + curr.first.substr(k+1, curr.first.size());
                if (has.find(s) != has.end()) {
                    has.erase(s);
                    vs v = curr.second;
                    v.push_back(s);
                    q.push({s, v});
                }
            }           
        }
    }
    

    return ans;
}

int main() {
    st start = "dog";
    st end = "cat";
    vs dic = {"dot", "dop", "dat", "cat"};

    vs ans =  stepword_chain(start, end, dic);
    for (auto s : ans) std::cout<<s<<", ";
    std::cout<<"\n";

    return 0;
}

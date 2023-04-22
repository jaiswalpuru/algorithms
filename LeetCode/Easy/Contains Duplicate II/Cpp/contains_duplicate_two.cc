#include <iostream>
#include <vector>
#include <map>

typedef std::vector<int> vi;
typedef std::map<int, int> mi;

bool contains_duplicate_two(vi &v, int k) {
    mi m;
    int n = v.size();
    for (int i=0;i<n;i++) {
        if (m.find(v[i]) != m.end()){
            int ind = m[v[i]];
            if (abs(i-ind) <= k) {
                return true;
            }else {
                m[v[i]] = i;
            }
        }else {
            m.insert(std::pair<int, int>(v[i], i));
        }
    }
    return false;
}

int main () {
    vi v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);
    v.push_back(1);
    std::cout<<contains_duplicate_two(v, 3)<<"\n";
} 

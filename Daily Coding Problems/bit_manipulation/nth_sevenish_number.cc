#include <iostream>
#include <set>
#include <vector>

typedef std::vector<int> vi;
typedef std::set<int> si;

si get_sevenish_number(int n) {
    vi powers;
    si totals;
    si tot; //acts as temp
            
    for (int i=0; i<n; i++) powers.push_back(std::pow(7, i));

    totals.insert(0);
    for (int num: powers) {
        for (int t : totals) {
            tot.insert(num+t);
        }
        totals.insert(tot.begin(), tot.end());
        tot.clear();
    }
    return totals;
}

int get_nth_seven_number(int n) {
    int i = 1;
    int cnt = 0, last_seven = 0;
    si s = get_sevenish_number(n);

    while (cnt < n) {
        if (s.find(i) != s.end()) {
            last_seven = i;
            cnt++;
        }
        i++;
    }
    return last_seven;

}

int get_nth_seven_number2(int n) {
    int bit_place = 0;
    int ans = 0;

    while(n) {
        if (n & 1)
            ans += std::pow(7, bit_place);
        n >>= 1;
        bit_place++;
    }
    return ans;
}

int main() {
    std::cout<<get_nth_seven_number2(7) << "\n";
    return 0;
}

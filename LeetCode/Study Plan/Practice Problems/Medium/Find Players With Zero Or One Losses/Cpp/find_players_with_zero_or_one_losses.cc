#include <iostream>
#include <vector>
#include <map>

#define PB push_back 

typedef std::vector<int> vi;
typedef std::vector<vi> vii;
typedef std::map<int, int> mi;

vii find_players_with_zero_or_one_losses(vii matches) {
    mi win, lost;
    int n = matches.size();
    for (int i=0; i<n; i++) {
        win[matches[i][0]]++;
        lost[matches[i][1]]++;
    }
    vi winners, just_winners;
    vii res;
    for (auto it= win.begin(); it!=win.end(); it++) if (lost.find(it->first) == lost.end()) winners.PB(it->first);
    for (auto it= lost.begin(); it!=lost.end(); it++) if (it->second == 1) just_winners.PB(it->first);
    res.PB(winners);
    res.PB(just_winners);
    return res;
}

int main() {
    vii res = find_players_with_zero_or_one_losses(vii{
        {{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}}
    });
    for (int i=0; i<res.size(); i++) {
        for (int j=0; j<res[i].size(); j++) std::cout<<res[i][j]<<" ";
        std::cout<<"\n";
    }
    std::cout<<"\n";
}

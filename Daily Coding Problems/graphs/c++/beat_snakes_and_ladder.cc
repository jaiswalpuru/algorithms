#include <iostream>
#include <map>
#include <vector>
#include <deque>
#include <utility>
#include <set>

typedef std::set<int> si;
typedef std::vector<int> vi;
typedef std::map<int, int> mii;
typedef std::deque<std::pair<int, int>> di;

int snakes_and_ladder(mii &snakes, mii &ladders) {
    int start = 0, end = 100;
    int turns = 0;
    vi v(101);
    si visited;
    di q;

    for (int i = 1; i <= 100; i++) v[i] = i;
    for (auto const &p : snakes)  v[p.first] = p.second;
    for (auto const &p : ladders) v[p.first] = p.second;

    q.push_back({start, turns});
    
    while (!q.empty()) {
        auto p = q.front(); // first is move , second is turns
        q.pop_front();

        for (int move = p.first + 1; move < p.first + 7; move++) {
            if (move >= end) return p.second + 1;
            if (visited.find(move) == visited.end()) {
                visited.insert(move);
                q.push_back({v[move], p.second + 1});
            }
        }
    }

    return -1;
}

int main() {
    mii snakes = {{17, 13}, {52, 29}, {57, 40}, {62, 22}, {88, 18}, {95, 51}, {97, 79}};
    mii ladders = {{3, 21}, {8, 30}, {28, 84}, {58, 77}, {75, 86}, {80, 100}, {90, 91}};
    std::cout<<snakes_and_ladder(snakes, ladders);

    return 0;
}

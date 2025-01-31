#include <iostream>

using namespace std;

int largest_overlap(vector<pair<int, int>> intervals) {
    vector<pair<int, char>> v;
    for (int i = 0; i < intervals.size(); i++) {
        v.push_back({intervals[i].first, 'S'});
        v.push_back({intervals[i].second, 'E'});
    }

    sort(v.begin(), v.end(), [&](pair<int, char>& a, pair<int, char>& b) {
            if (a.first == b.first) return a.second < b.second;
            return a.first < b.first;
            });
    
    int active = 0;
    int res = 0;
    for (int i = 0; i < v.size(); i++) {
        if (v[i].second == 'S') active++;
        else active--;
        res = max(res, active);
    }
    return res;
}


int main(int argc, char** argv) {
    cout << largest_overlap({{1, 3}, {2, 6}, {4, 8}, {6, 7}, {5, 7}});
    return 0;
}

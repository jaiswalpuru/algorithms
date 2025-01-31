#include <iostream>
#include <vector>

using namespace std;

vector<pair<int, int>> identify(vector<pair<int, int>>& int1, vector<pair<int, int>>& int2) {
    int i = 0;
    int j = 0;
    vector<pair<int, int>> res;

    while (i < int1.size() && j < int2.size()) {
        pair<int, int> a;
        pair<int, int> b;

        if (int1[i].first <= int2[j].first) {
            a = int1[i];
            b = int2[j];
        } else {
            a = int2[j];
            b = int1[i];
        }

        if (a.second >= b.first) {
            res.push_back({b.first, min(a.second, b.second)});
        }

        if (int1[i].second < int2[j].second) {
            i++;
        } else {
            j++;
        }
    }
    return res;
}

int main(int argc, char** argv) {
    vector<pair<int, int>> int1 = {{1, 4}, {5, 6}, {9, 10}};
    vector<pair<int, int>> int2 = {{2, 7}, {8, 9}};

    vector<pair<int, int>> res = identify(int1, int2);
    for (pair<int, int> p : res) cout << p.first << " " << p.second << "\n";
    return 0;
}

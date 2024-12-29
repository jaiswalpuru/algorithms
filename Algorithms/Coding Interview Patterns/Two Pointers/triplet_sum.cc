#include <iostream>

using namespace std;

void triplet_sum_brute_force(vector<int> v, vector<vector<int>> &res) {
    for (int i = 0; i < v.size(); i++) {
        for (int j = i + 1; j < v.size(); j++) {
            for (int k = j + 1; k < v.size(); k++) {
                if (v[i] + v[j] + v[k] == 0) res.push_back({v[i], v[j], v[k]});
            }
        }
    }
}

vector<vector<int>> pair_sum_sorted(vector<int> v, int start, int target);

void triplet_sum(vector<int> v, vector<vector<int>> &res) {
    sort(v.begin(), v.end());
    for (int i = 0; i < v.size(); i++) {
        if (v[i] > 0) break; // triplets consisting of only positive numbers will never equal to zero
        if (i > 0 && v[i] == v[i-1]) continue;
        vector<vector<int>> pairs = pair_sum_sorted(v, i + 1, -v[i]);
        for (int j = 0; j < pairs.size(); j++) {
            res.push_back({pairs[j][0], pairs[j][1], v[i]});
        }
    }

}

vector<vector<int>> pair_sum_sorted(vector<int> v, int start, int target) {
    vector<vector<int>> res;
    int i = start;
    int j = v.size() - 1;
    while (i < j) {
        int sum = v[i] + v[j];
        if (sum == target) {
            res.push_back({v[i], v[j]});
            i++;

            while (i < j && v[i] == v[i-1]) i++;
        } else if (sum < target) {
            i++;
        } else {
            j--;
        }
    }

    return res;
}

int main(int argc, char **argv) {
    vector<int> v = {0, -1, 2, -3, 1};
    vector<vector<int>> res;
    triplet_sum(v, res);
    
    for (int i = 0; i < res.size(); i++) {
        cout << res[i][0] << " " << res[i][1] << " " << res[i][2] << "\n";
    }
    
    return 0;
}


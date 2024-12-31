#include <iostream>
#include <unordered_map>

using namespace std;

vector<int> pair_sum_two_pass(vector<int> v, int target) {
    unordered_map<int, int> hash_map;

    for (int i = 0; i < v.size(); i++) hash_map[v[i]] = i;

    for (int i = 0; i < v.size(); i++) {
        int l = target - v[i];
        if (hash_map.find(l) != hash_map.end() && hash_map[l] != i) {
            return {l, v[i]};
        }
    }
    return {};
}

vector<int> pair_sum_one_pass(vector<int> v, int target) {
    unordered_map<int, int> hash_map;

    for (int i = 0; i < v.size(); i++) {
        int l = target - v[i];
        if (hash_map.find(l) != hash_map.end() && hash_map[l] != i) {
            return {l, v[i]};
        }
        hash_map[v[i]] = i;
    }

    return {};
}

int main(int argc, char **argv) {
    vector<int> res = pair_sum_one_pass({-1, 3, 4, 2}, 3);

    for (int i = 0; i < res.size(); i++) cout << res[i] << " ";
    cout << "\n";
    return 0;
}

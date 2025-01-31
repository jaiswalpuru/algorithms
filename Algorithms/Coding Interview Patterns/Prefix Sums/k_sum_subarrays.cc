#include <iostream>

using namespace std;

int k_sum_subarrays_hash_map(vector<int>& arr, int k) {
    unordered_map<int, int> umap;
    umap[0] = 1;

    int res = 0;
    int cur_sum = 0;
    for (int i = 0; i < arr.size(); i++) {
        cur_sum += arr[i];
        if (umap.count(cur_sum - k)) {
            res += umap[cur_sum - k];
        }
        umap[cur_sum]++;
    }
    return res;
}

int k_sum_subarrays_bf(vector<int>& arr, int k) {
    vector<int> prefix_sum = {0};
    for (int i = 0; i < arr.size(); i++) {
        prefix_sum.push_back(arr[i] + prefix_sum.back());
    }
    
    int cnt = 0;
    for (int j = 1; j < arr.size() + 1; j++) {
        for (int i = 1; i < j + 1; i++) {
            if (prefix_sum[j] - prefix_sum[i - 1] == k) {
                cnt++;
            }
        }
    }
    return cnt;
}

int main(int argc, char** argv) {
    vector<int> arr = {1, 2, -1, 1, 2};
    int k = 3;
    cout << k_sum_subarrays_hash_map(arr, k) << "\n";
    return 0;
}

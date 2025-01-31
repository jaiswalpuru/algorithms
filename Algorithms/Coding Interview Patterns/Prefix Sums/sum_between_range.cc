#include <iostream>

using namespace std;

vector<int> sum_range(vector<int>& arr, vector<vector<int>>& q) {
    vector<int> res;

    vector<int> prefix_sum = {arr[0]};
    for (int i = 1; i < arr.size(); i++) prefix_sum.push_back(prefix_sum[prefix_sum.size() - 1] + arr[i]);
    for (int i = 0; i < q.size(); i++) {
        if (q[i][0] == 0) {
            res.push_back(prefix_sum[q[i][1]]);
        } else {
            res.push_back(prefix_sum[q[i][1]] - prefix_sum[q[i][0] - 1]);
        }
    }
    return res;
}

int main(int argc, char** argv) {
    vector<int> arr = {3, -7, 6, 0, -2, 5};
    vector<vector<int>> q = {{0, 3}, {2, 4}, {2, 2}};
    vector<int> res = sum_range(arr, q);
    for (int i = 0; i < res.size(); i++) {
        cout << res[i] << " ";
    }
    cout << "\n";
    return 0;
}

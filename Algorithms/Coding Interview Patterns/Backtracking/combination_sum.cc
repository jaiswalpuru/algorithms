#include <iostream>
#include <vector>
#include <deque>

using namespace std;

void dfs(vector<int> arr, int sum, vector<vector<int>>& res, int curr_sum, deque<int> dq, int ind) {
    if (curr_sum == sum) {
        res.push_back({dq.begin(), dq.end()});
        return;
    }
    if (curr_sum > sum) return;

    for (int i = ind; i < arr.size(); i++) {
        dq.push_back(arr[i]);
        dfs(arr, sum, res, curr_sum + arr[i], dq, i);
        dq.pop_back();
    }
}


vector<vector<int>> combination_sum(vector<int> arr, int sum) {
    vector<vector<int>> res;
    dfs(arr, sum, res, 0, {}, 0);
    return res;
}

int main(int argc, char** argv) {
    vector<int> arr = {1, 2, 3};
    vector<vector<int>> res = combination_sum(arr, 4);
    for (auto vec : res) {
        for (auto val : vec) {
            cout << val << " ";
        }
        cout << "\n";
    }
    return 0;
}

#include <iostream>
#include <vector>
#include <deque>

using namespace std;

void _dfs(vector<int> nums, vector<vector<int>>& res, int i, deque<int> dq) {
    if (i == nums.size()) {
        res.push_back({dq.begin(), dq.end()});
        return;
    }

    dq.push_back(nums[i]);
    _dfs(nums, res, i + 1, dq);
    dq.pop_back();
    _dfs(nums, res, i + 1, dq);
}

vector<vector<int>> all_subsets(vector<int> nums) {
    vector<vector<int>> res;
    deque<int> dq;
    _dfs(nums, res, 0, dq);
    return res;
}

int main(int argc, char** argv) {
    vector<int> nums = {4, 5, 6};
    vector<vector<int>> res = all_subsets(nums);
    for (int i = 0; i < res.size(); i++) {
        for (int v : res[i]) {
            cout << v << " ";
        }
        cout << "\n";
    }
    return 0;
}

typedef vector<int> vi;
typedef vector<vi> vii;

class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vii res;
        backtrack(nums, res, nums.size(), 0);
        return res;
    }

    void backtrack(vi& nums, vii& res, int n, int ind) {
        if (ind == n) {
            res.push_back(nums);
            return;
        }

        for (int i=ind; i<n; i++) {
            swap(nums[i], nums[ind]);
            backtrack(nums, res, n, ind+1);
            swap(nums[i], nums[ind]);
        }
    }
};

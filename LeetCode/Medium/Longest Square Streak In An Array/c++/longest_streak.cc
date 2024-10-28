class Solution {
public:
    int longestSquareStreak(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        unordered_map<int, int> streak;
        int res = 0;

        for (int num : nums) {
            int root = (int)sqrt(num);

            if (root * root == num && streak.find(root) != streak.end()) {
                streak[num] = streak[root] + 1;
            } else {
                streak[num] = 1;
            }
        }

        for (auto& [k, v] : streak) {
            res = max(res, v);
        }

        return res == 1 ? -1 : res;
    }
};

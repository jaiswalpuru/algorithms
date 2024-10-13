class Solution {
public:
    vector<int> smallestRange(vector<vector<int>>& nums) {
        priority_queue<pair<int, pair<int, int>>, vector<pair<int, pair<int, int>>>, greater<>> pq;
        int max_val = INT_MIN;
        int range_start = 0;
        int range_end = INT_MAX;

        for (int i = 0; i < nums.size(); i++) {
            pq.push({nums[i][0], {i, 0}});
            max_val = max(max_val, nums[i][0]);
        }

        while (pq.size() == nums.size()) {
            auto [min_val, indices] = pq.top();
            pq.pop();
            int r = indices.first;
            int c = indices.second;

            if (max_val - min_val < range_end - range_start) {
                range_start = min_val;
                range_end = max_val;
            }
            
            // push the next val
            if (c + 1 < nums[r].size()) {
                pq.push({nums[r][c + 1], {r, c + 1}});
                max_val = max(max_val, nums[r][c + 1]);
            }
        }

        return {range_start, range_end};
    }
};

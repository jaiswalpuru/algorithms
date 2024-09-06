class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        vector<int> res;
        map<int, int> num_cnt;
        priority_queue<pair<int, int>> pq;

        for (int i = 0; i < nums.size(); i++) num_cnt[nums[i]]++;
        for (const auto& [key, val] : num_cnt) {
            pq.push(make_pair(val, key));
        }

        while (!pq.empty() && k > 0) {
            res.push_back(pq.top().second);
            pq.pop();
            k--;
        }

        return res;
    }
};

class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> hash_map;
        vector<int> res;
        for (int num : nums) hash_map[num]++;
        auto cmp = [](pair<int, int>& a, pair<int, int>& b) {
            return a.second > b.second;
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> pq;
        for (auto[key, val] : hash_map) {
            pq.push({key, val});
            if (pq.size() > k) pq.pop();
        }
        while (!pq.empty()) {
            res.push_back(pq.top().first);
            pq.pop();
        }
        sort(res.begin(), res.end());
        return res;
    }
};

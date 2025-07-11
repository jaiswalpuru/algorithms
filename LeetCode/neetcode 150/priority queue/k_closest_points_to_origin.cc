class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        vector<vector<int>> res;
        auto cmp = [&](pair<int, int>& a, pair<int, int>& b) {
            return get_dis(a.first, a.second) < get_dis(b.first, b.second);
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> pq(cmp);

        for (vector<int>& point : points) {
            pq.push({point[0], point[1]});
            if (pq.size() > k) pq.pop();
        }

        while (pq.size() > 0) {
            res.push_back({pq.top().first, pq.top().second});
            pq.pop();
        }

        return res;
    }

    int get_dis(int x, int y) {
        return x * x + y * y;
    }
};

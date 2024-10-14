class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>> res;

        sort(intervals.begin(), intervals.end());
        int start = intervals[0][0];
        int end =  intervals[0][1];
        for (int i = 0; i < intervals.size(); i++) {
            if (start == intervals[i][0] && end == intervals[i][1]) continue;

            if (intervals[i][0] <= end) {
                end = max(end, intervals[i][1]);
            } else {
                res.push_back({start, end});
                start = intervals[i][0];
                end = intervals[i][1];
            }
        }

        res.push_back({start, end});

        return res;
    }
};

class Solution {
public:
    static bool cmp(const vector<int>& a, const vector<int>& b) {
        if (a[0] == b[0]) return a[1] < b[1];
        return a[0] < b[0];
    }

    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        intervals.push_back(newInterval);
        sort(intervals.begin(), intervals.end(), cmp);
        int j=1;
        vector<vector<int>> res;
        vector<int> curr = intervals[0];
        while(j < intervals.size()) {
            if (curr[1] >= intervals[j][0]) {
                curr[1] = max(curr[1], intervals[j][1]);
            } else {
                res.push_back(curr);
                curr = intervals[j];
            }
            j++;
        }
        res.push_back(curr);
        return res;
    }
};

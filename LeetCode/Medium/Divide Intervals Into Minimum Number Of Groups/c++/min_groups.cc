class Solution {
public:
    int minGroups(vector<vector<int>>& intervals) {
        int concurr_in = 0;
        int max_concurr = 0;
        vector<pair<int, int>> intervals_end;

        for (auto& in : intervals) {
            intervals_end.push_back({in[0], 1});
            intervals_end.push_back({in[1]+1, -1});
        }

        sort(intervals_end.begin(), intervals_end.end());

        for (auto in : intervals_end) {
            concurr_in += in.second;
            max_concurr = max(max_concurr, concurr_in);
        }

        return max_concurr;
    }
};

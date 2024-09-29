class Solution {
public:
    int maxDistance(vector<vector<int>>& arrays) {
        int min_val = INT_MAX;
        int max_val = INT_MIN;
        int res = INT_MIN;
        
        for (int num : arrays[0]) {
            min_val = min(min_val, num);
            max_val = max(max_val, num);
        }
        
        for (int i = 1; i < arrays.size(); i++) {
            int t_min = INT_MAX;
            int t_max = INT_MIN;
            for (int num : arrays[i]) {
                t_min = min(t_min, num);
                t_max = max(t_max, num);
            }
            res = max(res, max(abs(min_val - t_max), abs(max_val - t_min)));
            min_val = min(min_val, t_min);
            max_val = max(max_val, t_max);
        }

        return res;
    }
};

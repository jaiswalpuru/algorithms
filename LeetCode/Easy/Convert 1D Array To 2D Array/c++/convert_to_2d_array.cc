class Solution {
public:
    vector<vector<int>> construct2DArray(vector<int>& original, int m, int n) {
        vector<vector<int>> res;
        int k = 0;

        if (m * n != original.size()) return res;

        res.resize(m);
        for (int i = 0; i < m; i++) {
            res[i].resize(n);
            for (int j = 0; j < n; j++) {
                res[i][j] = original[k++];
            }

        }

        return res;
    }
};

class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int> res;
        int row = matrix.size();
        int col = matrix[0].size();
        int left = 0;
        int right = col-1;
        int up = 0;
        int down = row-1;

        while (res.size() < row * col) {
            for (int i = left; i <= right; i++) res.push_back(matrix[up][i]);
            for (int i = up + 1; i <= down; i++) res.push_back(matrix[i][right]);

            if (up != down)
                for (int i = right - 1; i >= left; i--) res.push_back(matrix[down][i]);
            if (left != right)
                for (int i = down - 1; i > up; i--) res.push_back(matrix[i][left]);

            left++;
            right--;
            up++;
            down--;
        }

        return res;
    }
};

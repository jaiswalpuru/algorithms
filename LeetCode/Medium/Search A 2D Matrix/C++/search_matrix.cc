class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int size = matrix.size();
        for (int i=0; i<size; i++)
            if (binary_search(matrix[i].begin(), matrix[i].end(), target)) return true;
        return false;
    }
};

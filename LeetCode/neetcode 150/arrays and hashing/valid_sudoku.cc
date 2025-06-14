class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        vector<unordered_set<int>> rows(9);
        vector<unordered_set<int>> cols(9);
        vector<vector<unordered_set<int>>> sub_grid(9, vector<unordered_set<int>>(9));
        for (int r = 0; r < board.size(); r++) {
            for (int c = 0; c < board[0].size(); c++) {
                if (board[r][c] == '.') continue;
                int num = board[r][c] - '0';
                if (rows[r].count(num)) return false;
                if (cols[c].count(num)) return false;
                if (sub_grid[r / 3][c / 3].count(num)) return false;
                rows[r].insert(num);
                cols[c].insert(num);
                sub_grid[r / 3][c / 3].insert(num);
            }
        }
        return true;
    }
};

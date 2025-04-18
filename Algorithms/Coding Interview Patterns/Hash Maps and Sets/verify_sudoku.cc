#include <iostream>
#include <unordered_map>

using namespace std;

bool valid_sudoku(vector<vector<char>>& mat) {
    vector<unordered_set<int>> row(9);
    vector<unordered_set<int>> col(9);
    vector<vector<unordered_set<int>>> sub_grid(9, vector<unordered_set<int>>(9));
    for (int r = 0; r < board.size(); r++) {
        for (int c = 0; c < board[0].size(); c++) {
            if (board[r][c] == '.') continue;
            int num = board[r][c] - '0';
            if (row[r].count(num)) return false;
            if (col[c].count(num)) return false;
            if (sub_grid[r / 3][c / 3].count(num)) return false;
            row[r].insert(num);
            col[c].insert(num);
            sub_grid[r / 3][c / 3].insert(num);
        }
    }
    return true;
}

int main(int argc, char** argv) {
    
    return 0;
}

#include <iostream>
#include <vector>

using namespace std;

class Solution{
    public:
        void solveSudoku(vector<vector<char> > &board){
            sudoku_solver(board);
        }

        bool sudoku_solver(vector<vector<char> > &board);
        bool is_safe(vector<vector<char> > &board, int row, int col, char c);
};

bool Solution::sudoku_solver(vector<vector<char> > &board){
    for (int i=0; i<9; i++){
        for (int j=0; j<9; j++){
            if (board[i][j] == '.'){
                for (int k='1';k<='9'; k++){
                    if (is_safe(board,i,j,k)){
                        board[i][j] = k;
                        if (sudoku_solver(board)) {
                            return true;
                        }else{
                            board[i][j] = '.';
                        }
                    }
                }
                return false;
            }
        }
    }
    return true;
}

bool Solution::is_safe(vector<vector<char> > &board, int row, int col, char c){
    for (int i=0; i<9; i++) {
        if (board[i][col] == c){
            return false;
        }
        if (board[row][i] == c){
            return false;
        }

        if (board[3*(row/3) + i/3][3*(col/3)+i%3]==c){
            return false;
        }
    }
    return true;
}

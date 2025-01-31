#include <iostream>

using namespace std;

bool binary_search(vector<vector<int>> a, int r, int target) {
    int l = 0;
    int rr = a[r].size() - 1;

    while (l <= rr) {
        int mid = l + (rr - l) / 2;
        if (a[r][mid] == target) return true;
        else if (a[r][mid] > target) {
            rr = mid - 1;
        } else {
            l = mid + 1;
        }
    }
    return false;
}

bool mat_search1(vector<vector<int>> mat, int target) {
    for (int i = 0; i < mat.size(); i++) {
        if (binary_search(mat, i, target)) return true;
    }
    return false;
}

bool mat_search2(vector<vector<int>> mat, int target) {
    int m = mat.size();
    int n = mat[0].size();
    int l = 0;
    int r = m * n - 1;

    while(l <= r) {
        int mid = l + (r - l) / 2;
        int rr = mid / n;
        int cc = mid % n;
        if (mat[rr][cc] == target) return true;
        if (mat[rr][cc] > target) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    return false;
}


int main(int argc, char** argv) {
    vector<vector<int>> mat = {{2, 3, 4, 6}, {7, 10, 11, 17}, {20, 21, 24, 33}};
    cout << mat_search2(mat, 21) << "\n";
    cout << mat_search2(mat, 99) << "\n";

    return 0;
}

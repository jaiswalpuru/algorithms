#include <iostream>
#include <vector> 

using namespace std;

 vector<vector<int>> merge_intervals(vector<vector<int>> arr) {
     sort(arr.begin(), arr.end(), [&](vector<int>& a, vector<int>& b) {
             return a[0] < b[0];
             });
     vector<vector<int>> res;
     res.push_back(arr[0]);
     
     for (int i = 1; i < arr.size(); i++) {
        vector<int> b = res.back();
        if (b[1] < arr[i][0]) {
            res.push_back(arr[i]);
        } else {
            b[1] = max(b[1], arr[i][1]);
            res[res.size() - 1][1] = b[1];
        }
     }
     return res;
 }

int main(int argc, char** argv) {
    vector<vector<int>> arr = {{3, 4}, {7, 8}, {2, 5}, {6, 7}, {1, 4}};
    vector<vector<int>> res = merge_intervals(arr);
    for (vector<int> v : res) cout << v[0] << " " << v[1] << "\n";
    return 0;
}

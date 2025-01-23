#include <iostream>
#include <unordered_map>
#include <deque>

using namespace std;

void _dfs(vector<int> arr, unordered_map<int, int>& visited, vector<vector<int>>& res, deque<int> s) {
    if (s.size() == arr.size()) {
        res.push_back({s.begin(), s.end()});
        return;
    }
    
    for (int i = 0; i < arr.size(); i++) {
        if (!visited[arr[i]]) {
            s.push_back(arr[i]);
            visited[arr[i]] = true;
            _dfs(arr, visited, res, s);
            visited[arr[i]] = false;
            s.pop_back();
        }
    }
}

vector<vector<int>> all_permutation(vector<int> arr) {
    vector<vector<int>> res;
    deque<int> dq;
    unordered_map<int, int> visited;
    _dfs(arr, visited, res, dq);
    return res;
}

int main(int argc, char** argv) {
    vector<int> arr = {4, 5, 6};
    vector<vector<int>> res = all_permutation(arr);
    for (int i = 0; i < res.size(); i++) {
        cout << res[i][0] << " " << res[i][1] << " " << res[i][2] << "\n";
    }
    return 0;
}

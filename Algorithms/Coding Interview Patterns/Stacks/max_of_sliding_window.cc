#include <iostream>
#include <deque>

using namespace std;

vector<int> max_of_sliding_window(vector<int>& arr, int k) {
    deque<pair<int, int>> dq;
    int l = 0;
    int r = 0;
    vector<int> res;

    while (r < arr.size()) {
        while (!dq.empty() && arr[r] >= dq.back().first) {
            dq.pop_back();
        }
        
        dq.push_back({arr[r], r});
        if (r - l + 1 == k) {
            if (dq.front().second < l) dq.pop_front();
            res.push_back(dq.front().first);
            l++;
        }
        r++;
    }
    return res;
}

int main(int argc, char** argv) {
    vector<int> arr = {3, 2, 4, 1, 2, 1, 1};
    int k = 4;
    vector<int> res = max_of_sliding_window(arr, k);
    for (int v : res) cout << v << " ";
    cout << "\n";
    return 0;
}

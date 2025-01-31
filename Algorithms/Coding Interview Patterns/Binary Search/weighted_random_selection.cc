#include <iostream>

using namespace std;

int weighted_index_pick(vector<int>& arr) {
    vector<int> prefix_sum = {arr[0]};
    for (int i = 1; i < arr.size(); i++) prefix_sum.push_back(prefix_sum[prefix_sum.size() - 1] + arr[i]);

    //select
    //
    int target = 1 + rand() % (prefix_sum[prefix_sum.size() - 1] - 1);
    int l = 0;
    int r = prefix_sum.size() - 1;

    while (l < r) {
        int mid = l + (r - l) / 2;
        if (prefix_sum[mid] < target) {
            l = mid + 1;
        } else {
            r = mid;
        }
    }
    cout << target << " : ";
    return l;
}


int main(int argc, char** argv) {
    vector<int> arr = {3, 1, 2, 4};
    cout << weighted_index_pick(arr) << "\n";
    return 0;
}

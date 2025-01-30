#include <iostream>

using namespace std;

int binary_srch(vector<int>& arr, int val) {
    int l = 0;
    int r = arr.size() - 1;

    while(l <= r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] == val) return mid;
        if (arr[mid] > val) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
    return -1;
}

int main(int argc, char** argv) {
    vector<int> arr = {1, 3, 4, 9, 10, 12, 14, 17, 20};
    cout << binary_srch(arr, 3) << "\n";
    return 0;
}

#include <iostream>

using namespace std;

int lower_b(vector<int>& arr, int val) {
    int l = 0;
    int r = arr.size() - 1;

    while (l < r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] > val) {
            r = mid - 1;
        } else if (arr[mid] < val) {
            l = mid + 1;
        } else {
            r = mid;
        }
    }
    return val == arr[l] ? l : -1;
}

int upper_b(vector<int>& arr, int val) {
    int l = 0;
    int r = arr.size() - 1;
    while (l < r) {
        int mid = l + (r - l) / 2 + 1;
        if (arr[mid] > val) {
            r = mid - 1;
        } else if (arr[mid] < val) {
            l = mid + 1;
        } else {
            l = mid;
        }
    }
    return arr[r] == val ? r : -1;
}

vector<int> first_and_last(vector<int> arr, int val) {
    return {lower_b(arr, val), upper_b(arr, val)};
}

int main(int argc, char** argv) {
    vector<int> arr = {1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 9, 10};
    vector<int> res = first_and_last(arr, 4);
    cout << res[0] << " " << res[1] << "\n";
    return 0;
}

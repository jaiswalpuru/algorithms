#include <iostream>

using namespace std;

int find_target(vector<int> arr, int target) {
    int l = 0;
    int r = arr.size() - 1;

    while (l < r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] == target) {
            return mid;
        } else if (arr[l] <= arr[mid]) {
            if (arr[l] <= target && target < arr[mid]) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        } else { 
            if (arr[mid] < target && target <= arr[r]) {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
    }

    return arr[l] == target ? l : -1;
}

int main(int argc, char** argv) {
    cout <<find_target({8, 9, 1, 2, 3, 4, 5, 6, 7}, 1) << "\n";
    return 0;
}

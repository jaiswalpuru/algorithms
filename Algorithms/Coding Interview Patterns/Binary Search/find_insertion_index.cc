#include <iostream>

using namespace std;

int get_insertion_index(vector<int>& arr, int target) {
    int l = 0;
    int r = arr.size();

    while (l < r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] >= target) r = mid;
        else l = mid + 1;
    }
    return l;
}

int main(int argc, char** argv) {
    vector<int> arr = {1, 2, 4, 5, 7, 8, 9};
    int target = 6;

    cout << get_insertion_index(arr, target) << "\n";
    return 0;
}

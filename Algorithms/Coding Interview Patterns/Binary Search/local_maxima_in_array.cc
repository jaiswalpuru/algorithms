#include <iostream>

using namespace std;

int local_maxima(vector<int> arr) {
    int l = 0;
    int r = arr.size() - 1;

    while (l < r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] > arr[mid + 1]) {
            r = mid;
        } else {
            l = mid + 1;
        }
    }
    return l;
}

int main(int argc, char** argv) {
    cout << local_maxima({1, 4, 3, 2, 3}) << "\n";
    return 0;
}

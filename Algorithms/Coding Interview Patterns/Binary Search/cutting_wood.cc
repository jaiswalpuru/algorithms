#include <iostream>

// note : always bias the midpoint to the right during upper bound search

using namespace std;

int can_cut(int h, int k, vector<int> arr) {
    int wood = 0;
    for (int v : arr) {
        if (v > h) wood += (v - h);
    }
    return wood >= k;
}

int cutting_wood(vector<int> arr, int k) {
    int l = 0;
    int r = *max_element(arr.begin(), arr.end());

    while (l < r) {
        int height = l + (r - l) / 2 + 1;
        if (can_cut(height, k, arr)) {
            l = height;
        } else {
            r = height - 1;
        }
    }
    return r;
}

int main(int argc, char** argv) {
    cout << cutting_wood({2, 6, 3, 8}, 7) << "\n";
    return 0;
}

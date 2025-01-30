#include <iostream>
#include <vector>

using namespace std;

void sort(vector<int>& arr) {
    int left = 0;
    int right = arr.size() - 1;

    int i = 0;
    while (i <= right) {
        if (arr[i] == 0) {
            swap(arr[i], arr[left]);
            left++;
            i++;
        } else if (arr[i] == 2) {
            swap(arr[i], arr[right]);
            right--;
        } else {
            i++;
        }
    }
}

int main(int argc, char** argv) {
    vector<int> arr = {0, 1, 2, 0, 1, 2, 0};
    sort(arr);
    for (int v : arr) cout << v << " ";
    cout << "\n";
    return 0;
}

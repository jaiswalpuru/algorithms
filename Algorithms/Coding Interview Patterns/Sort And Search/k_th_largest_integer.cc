#include <iostream>
#include <queue>

using namespace std;

int partition(vector<int>& arr, int left, int right) {
    int pivot = arr[right];
    int lo = left;
    for (int i = left; i < right; i++) {
        if (arr[i] < pivot) {
            swap(arr[i], arr[lo]);
            lo++;
        }
    }
    swap(arr[lo], arr[right]);
    return lo;
}

int helper(vector<int>& arr, int left, int right, int k) {
    int n = arr.size();
    if (left >= right) arr[left];

    int random_index = left + rand() % (right - left + 1);
    swap(arr[random_index], arr[right]);
    int pivot_index = partition(arr, left, right);
    if (pivot_index < n - k) return helper(arr, pivot_index + 1, right, k);
    else if (pivot_index > n - k) return helper(arr, left, pivot_index - 1, k);
    else return arr[pivot_index];
}

int k_thlargest_using_quickselect(vector<int> arr, int k) {
    return helper(arr, 0, arr.size() - 1, k);
}

int k_thlargest_using_heap(vector<int> nums, int k) {
    priority_queue<int, vector<int>, greater<int>> pq;

    for (int n : nums) {
        pq.push(n);
        if (pq.size() > k) pq.pop();
    }
    return pq.top();
}

int main(int argc, char** argv) {
    vector<int> nums = {5, 2, 4, 3, 1, 6};
    int k = 3;
    cout << k_thlargest_using_quickselect(nums, k) << "\n";
    return 0;
}

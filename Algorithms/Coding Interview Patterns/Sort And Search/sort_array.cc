#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// for practice
// quicksort
// merge sort


void merge(vector<int>& arr, int left, int mid, int right) {
    int n1 = mid - left + 1;
    int n2 = right - mid;
    vector<int> l(n1), r(n2);

    for (int i = 0; i < n1; i++) l[i] = arr[left + i];
    for(int i = 0; i < n2; i++) r[i] = arr[mid + 1 + i];

    int i = 0, j = 0;
    int k = left;

    while (i < n1 && j < n2) {
        if (l[i] <= r[j]) {
            arr[k] = l[i];
            i++;
        } else {
            arr[k] = r[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        arr[k] = l[i];
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = r[j];
        j++;
        k++;
    }
}

void merge_sort_helper(vector<int>& arr, int left, int right) {
    if (left >= right) return;
    int mid = left + (right - left) / 2;
    merge_sort_helper(arr, left, mid);
    merge_sort_helper(arr, mid + 1, right);
    merge(arr, left, mid, right);
}

void merge_sort(vector<int>& arr) {
    merge_sort_helper(arr, 0, arr.size() - 1);
}

// couting sort the values in the array do not contain negative values

void counting_sort(vector<int> arr) {
    vector<int> res;
    auto max_val = max_element(arr.begin(), arr.end());
    vector<int> count(*max_val + 1, 0);

    for (int val : arr) {
        count[val]++;
    }

    for (int i = 0; i < count.size(); i++) {
        for (int j = count[i]; j > 0; j--)
            res.push_back(i);
    }
    cout << "couting sort \n";
    for (int v : res) cout << v << " ";
    cout << "\n";
}

int partition(vector<int>& arr, int left, int right) {
    int pivot = arr[right];
    int lo = left;

    for (int i = left; i <= right; i++) {
        if (arr[i] < pivot) {
            swap(arr[i], arr[lo]);
            lo++;
        }
    }
    swap(arr[right], arr[lo]);
    return lo;
}

void quicksort_helper(vector<int>& arr, int left, int right) {
    if (left >= right) return;

    int pivot = partition(arr, left, right);
    quicksort_helper(arr, left, pivot - 1);
    quicksort_helper(arr, pivot + 1, right);
}

void quicksort(vector<int>& arr) {
    quicksort_helper(arr, 0, arr.size() - 1);
}

int main(int argc, char** argv) {
    vector<int> arr = {6, 8, 4, 2, 7, 3, 1, 5};
    //counting_sort(arr);
    //quicksort(arr);
    vector<int> a = {38, 27};
    merge_sort(a);
    for (int val : a) cout << val << " ";
    cout << "\n";
    return 0;
}

#include <iostream>
#include <queue>

using namespace std;

void sort_k_sorted(vector<int>& arr, int k) {
    priority_queue<int, vector<int>, greater<int>> pq;
    int ind = 0;

    for (int i = 0; i <= k; i++) {
        pq.push(arr[i]);
    }

    for (int i = k + 1; i < arr.size(); i++) {
        arr[ind] = pq.top();
        pq.pop();
        pq.push(arr[i]);
        ind++;
    }

    while(!pq.empty()) {
        arr[ind] = pq.top();
        pq.pop();
        ind++;
    }
}

int main(int argc, char** argv) {
    vector<int> arr = {5, 1, 9, 4, 7, 10};
    for (int v : arr) cout << v << " ";
    cout << "\n";
    int k = 2;
    sort_k_sorted(arr, k);
    for (int v : arr) cout << v << " ";
    cout << "\n";
    return 0;
}

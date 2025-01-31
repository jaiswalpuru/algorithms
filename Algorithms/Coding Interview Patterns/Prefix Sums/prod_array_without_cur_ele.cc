#include <iostream>

using namespace std;

//brute would be to multiply all the elements and then for each
// place divide the prod with the current element

vector<int> prod(vector<int> arr) {
    vector<int> res(arr.size());

    res[0] = 1;
    for (int i = 1; i < arr.size(); i++) {
        res[i] = arr[i - 1] * res[i - 1];
    }

    int right = 1;
    for (int i = arr.size() - 1; i >= 0; i--) {
        res[i] *= right;
        right *= arr[i];
    }
    return res;
}

    

int main(int argc, char** argv) {
    vector<int> res = prod({2, 3, 1, 4, 5});
    for (int v : res) {
        cout << v << " ";
    }
    cout << "\n";
    return 0;
}

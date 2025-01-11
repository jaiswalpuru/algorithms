#include <iostream>
#include <vector>

using namespace std;

void shift_zero_method_one(vector<int>& v) {
    vector<int> res(v.size());
    int j = 0;

    for (int i = 0; i < v.size(); i++) {
        if (v[i] != 0) {
            res[j] = v[i];
            j++;
        }
    }
    
    v = res;
}

void shift_zero_efficient(vector<int> &v) {
    int l = 0;
    int r = v.size() - 1;

    while (l < r) {
        while (v[r] == 0 && r > l) {
            r--;
        }
        if (v[l] == 0) {
            int temp = v[l];
            v[l] = v[r];
            v[r] = temp;
        }
        l++;
    }
}

int main(int argc, char **argv) {
    vector<int> v = {0, 1, 0, 3, 2};
    shift_zero_efficient(v);
    for (auto val : v) cout << val << " ";
    cout << "\n";
    return 0;
}

#include <iostream>

using namespace std;

void two_sum_sorted_brute(vector<int> v, int target, vector<int> &res) {
    for (int i = 0; i < v.size(); i++) {
        for (int j = i + 1; j < v.size(); j++) {
            if (v[i] + v[j] == target) {
                res.push_back(v[i]);
                res.push_back(v[j]);
            }
        }
    }
    return;
}

void two_sum_sorted(vector<int> v, int target, vector<int> &res) {
    int i = 0;
    int j = v.size() - 1;
    
    while (i < j) {
        if (v[i] + v[j] == target) {
            res.push_back(v[i]);
            res.push_back(v[j]);
            break;
        } else if (v[i] + v[j] < target) {
            i++;
        } else {
            j--;
        }
    }

}

int main(int argc, char **argv) {
    vector<int> v = {-5, -2, 3, 4, 6};
    vector<int> res;
    two_sum_sorted(v, 7, res);

    for (int i = 0; i < res.size(); i++) {
        cout<<res[i]<<" ";
    }
    cout<<"\n";

    return 0;
}


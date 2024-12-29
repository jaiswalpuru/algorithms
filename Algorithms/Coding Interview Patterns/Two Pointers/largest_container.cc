#include <iostream>

using namespace std;

int largest_container_brute_force(vector<int> v) {
    int max_water = 0;
    for (int i = 0; i < v.size(); i++) {
        for (int  j = i + 1; j < v.size(); j++) {
            max_water = max(max_water, min(v[i], v[j]) * (j - i));
        }
    }

    return max_water;
}

int largest_container(vector<int> v) {
    int max_water = 0;

    int i = 0;
    int j = v.size() - 1;

    while (i < j) {
        max_water = max(max_water, min(v[i], v[j]) * (j - i));
        if (v[i] < v[j]) {
            i++;
        } else if (v[i] > v[j]) {
            j--;
        } else {
            i++;
            j--;
        }
    }

    return max_water;
}

int main(int argc, char **argv) {
    cout << largest_container({2, 7, 8, 3, 7, 6}) << "\n";
    return 0;
}


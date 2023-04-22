#include <iostream>
#include <vector>

typedef std::vector<int> vi;

int remove_element(vi v, int val) {
    int j = 0;
    for (int i=0; i<v.size(); i++) {
        if (v[i] != val) {
            v[j] = v[i];
            j++;
        }
    }
    return j;
}

int main() {
    std::cout<<remove_element(vi{3,2,2,3}, 3)<<"\n";
}

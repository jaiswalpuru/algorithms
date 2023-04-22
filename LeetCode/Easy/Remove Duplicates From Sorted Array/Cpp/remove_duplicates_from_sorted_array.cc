#include <iostream>
#include <vector>

typedef std::vector<int> vi;

int remove_duplicates_from_sorted_array(vi v) {
    int i=0, j=0;
    while(i<v.size() && j<v.size()) {
        int curr = v[j];
        while(j<v.size() && v[j] == curr) {
            j++;
        }
        v[i] = v[j-1];
        i++;
    }
    return i;
}

int main() {
    std::cout<<remove_duplicates_from_sorted_array(vi{1,1,2})<<"\n";
}

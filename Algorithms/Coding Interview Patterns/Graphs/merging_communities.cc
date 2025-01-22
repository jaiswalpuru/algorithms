#include <iostream>
#include <vector>

using namespace std;

int find(int x, vector<int>& parent) {
    if (x == parent[x]) {
        return x;
    }
    parent[x] = find(parent[x], parent);
    return parent[x];
}

void union_of_two_nodes(int x, int y, vector<int>& parent, vector<int>& size) {
    x = find(x, parent);
    y = find(y, parent);

    if (x == y) return;
    if (size[x] >= size[y]) {
        parent[y] = x;
        size[x] += size[y];
    } else {
        parent[x] = y;
        size[y] += size[x];
    }
}

int main(int argc, char** argv) {
    vector<int> parent(5);
    vector<int> size(5, 1);

    for (int i = 0; i < 5; i++) parent[i] = i;
    union_of_two_nodes(0, 1, parent, size);
    union_of_two_nodes(1, 2, parent, size);
    cout << "size of 3 : " << size[3] << "\n";
    cout << "size of 0 : " << size[0] << "\n";
    union_of_two_nodes(3, 4, parent, size);
    cout << "size of 3 : " << size[3] << "\n";
    return 0;
}

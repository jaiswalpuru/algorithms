#include <iostream>
#include <vector>

using namespace std;

int find(int x, vector<int>& parent) {
    if (x == parent[x]) return x;
    parent[x] = find(parent[x], parent);
    return parent[x];
}

bool union_node(int x, int y, vector<int>& parent, vector<int>& size) {
    x = find(x, parent);
    y = find(y, parent);
    if (x == y) return false;
    
    if (size[x] > size[y]) {
        size[x] += size[y];
        parent[y] = x;
    } else {
        size[y] += size[x];
        parent[x] = y;
    }
    return true;
}

int connect_dots(vector<vector<int>> points) {
    vector<vector<int>> edges;
    int n = points.size();
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            int cost = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
            edges.push_back({cost, i, j});
        }
    }

    sort(edges.begin(), edges.end());
    vector<int> parent(n);
    for (int i = 0; i < n; i++) parent[i] = i;
    vector<int> size(n, 1);
    int tot_cost = 0;
    int edges_added = 0;
    for (auto curr : edges) {
        if (union_node(curr[1], curr[2], parent, size)) {
            tot_cost += curr[0];
            edges_added++;
            if (edges_added == n - 1) {
                return tot_cost;
            }
        }
    }
    return -1;
}

int main(int argc, char** argv) {
    vector<vector<int>> points = {{1, 1}, {2, 6}, {3, 2}, {4, 3}, {7, 1}};
    int min_cost = connect_dots(points);
    cout << min_cost << "\n";
    return 0;
}

#include <iostream>
#include <vector>

using namespace std;

bool dfs(int node, int color, vector<vector<int>> g, vector<int>& colors) {
    colors[node] = color;
    for (int j = 0; j < g[node].size(); j++) {
        if (colors[g[node][j]] == color) return false;
        if (colors[g[node][j]] == 0 && !dfs(g[node][j], -color, g, colors)) return false;
    }
    return true;
}

bool is_graph_bipartite(vector<vector<int>> g, int nodes) {
    vector<int> colors(nodes, 0);
    for (int i = 0; i <nodes; i++) {
        if (colors[i] == 0 && !dfs(i, 1, g, colors)) return false;
    }
    return true;
}

int main(int argc, char** argv) {
    vector<vector<int>> g = {{1, 4}, {0, 2}, {1}, {4}, {0, 3}};
    cout << is_graph_bipartite(g, 5) << "\n";
    return 0;
}

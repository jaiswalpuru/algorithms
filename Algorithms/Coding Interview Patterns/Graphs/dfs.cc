#include <iostream>
#include <set>

using namespace std;

void _dfs(int start, vector<vector<int>> graph, set<int>& visited) {
    if (visited.find(start) != visited.end()) return;

    visited.insert(start);
    for (int i = 0; i < graph[start].size(); i++) _dfs(graph[start][i], graph, visited);
}

 void dfs(vector<vector<int>> graph) {
    set<int> visited;

    for (int i = 0; i < 5; i++) {
        _dfs(i, graph, visited);
    }

    for (auto v : visited) {
        cout << v << " ";
    }
    cout << "\n";
 }

int main(int argc, char** argv) {
    vector<vector<int>> graph = {{1, 2, 3}, {0}, {3, 4, 0}, {0, 2}, {2}};
    dfs(graph);
    return 0;
}

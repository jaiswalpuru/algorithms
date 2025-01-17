#include <iostream>
#include <vector>
#include <set>
#include <queue>

using namespace std;

void bfs(vector<vector<int>> graph) {
    set<int> visited;
    queue<int> q;

    q.push(0);

    while(!q.empty()) {
        int curr = q.front();
        q.pop();
        if (visited.find(curr) != visited.end()) continue;
        visited.insert(curr);
        for (int i = 0; i < graph[curr].size(); i++) {
            if (visited.find(graph[curr][i]) != visited.end()) continue;
            q.push(graph[curr][i]);
        }
    }

    for (auto v : visited) cout << v << " " << "\n";
}

int main(int argc, char **argv) {
    vector<vector<int>> graph = {{1, 2, 3}, {0}, {0, 3, 4}, {0, 2}, {2}};
    bfs(graph);
    return 0;
}

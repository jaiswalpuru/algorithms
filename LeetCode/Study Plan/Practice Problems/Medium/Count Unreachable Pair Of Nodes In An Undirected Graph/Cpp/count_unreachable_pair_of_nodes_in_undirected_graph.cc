#include <iostream>
#include <vector>
#include <unordered_map>

typedef std::vector<int> vi;
typedef std::vector<vi> vii;
typedef std::unordered_map<int, vi> umi;
typedef std::vector<bool> vb;
typedef long long ll;

umi make_graph(vii &edges) {
    umi g;
    for (int i=0;i<edges.size(); i++) {
        g[edges[i][0]].push_back(edges[i][1]);
        g[edges[i][1]].push_back(edges[i][0]);
    }
    return g;
}

void dfs(int start, umi &graph, vb &visited, ll &cnt) {
    visited[start]=true;
    cnt++;
    for (int i=0;i<graph[start].size(); i++)
        if (!visited[graph[start][i]]) dfs(graph[start][i], graph, visited, cnt);
}

ll count_unreachable_pair_of_nodes_in_undirected_graph(int n, vii &edges) {
    umi graph = make_graph(edges);
    vb visited(n);
    ll total = ((ll)n*(n-1))/2;
    for (int i=0;i<n;i++) {
        ll cnt = 0;
        if (!visited[i]) {
            dfs(i, graph, visited, cnt);
            total -= (cnt*(cnt-1))/2;
        }
    }
    return total;
}


int main() {
    vii edges{{0,1},{0,2}, {1,2}};
    std::cout<<count_unreachable_pair_of_nodes_in_undirected_graph(3, edges)<<"\n";
}

#include <iostream>
#include <vector>
#include <map>

typedef std::vector<int> vi;
typedef std::map<int, vi> mvi;
typedef std::map<int, int> mi;

int dfs(mvi g, int start, mi &desc) {
    int d = 0;

    for (auto nei : g[start]) {
        int nodes = dfs(g, nei, desc);

        desc[nei] += nodes-1;
        d += nodes;
    }
    return d + 1;
}

int max_len(mvi g, int start) {
    mi desc;
    vi v;

    for (auto const &p : g) desc[p.first] = 0;

    dfs(g, start, desc);

    for (auto const &p : desc) if (p.second % 2 == 1) {
        v.push_back(p.second);
    }
    return v.size();
}

int main() {
    mvi graph;
    
    graph[1] = {2, 3};
    graph[2] = {};
    graph[3] = {4, 5};
    graph[4] = {6, 7, 8};
    graph[5] = {};
    graph[6] = {};
    graph[7] = {};
    graph[8] = {};

    std::cout<<max_len(graph, 1)<<"\n";
}

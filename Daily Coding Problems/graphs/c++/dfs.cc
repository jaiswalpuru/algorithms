#include <iostream>
#include <string>
#include <list>
#include <map>

typedef std::string ss;
typedef std::list<ss> ls;
typedef std::map<ss, ls> mls;
typedef std::map<ss, bool> mlb;

void dfs(mls graph, ss start, mlb& visited) {
    visited[start] = true;
    
    for (auto nei : graph[start]) {
        if (visited.find(nei) == visited.end()) {
            dfs(graph, nei, visited);
        }
    }
}

int main() {
    mls graph;
    mlb visited;
    graph["JFK"] = {"SFO", "LAX"};
    graph["SFO"] = {"ORL"};
    graph["ORL"] = {"JFK", "LAX", "DFW"};
    graph["LAX"] = {"DFW"};

    dfs(graph, "ORL",  visited);
    for (auto p : visited) {
        std::cout<<p.first<<" "<<p.second<<"\n";
    }
}

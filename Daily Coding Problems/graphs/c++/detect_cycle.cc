// given an undirected graph determine if there exists a cycle or not

#include <iostream>
#include <vector>
#include <map>

typedef std::vector<int> vi;
typedef std::map<int, vi> mi;
typedef std::map<int, bool> mb;

bool search(mi graph, mb &visited, int start, int parent) {
    visited[start] = true;
    
    for (auto nei : graph[start]) {
        if (!visited[nei]) {
            if (search(graph, visited, nei, start)) return true;
        } else if (parent != nei) return true;
    }
    
    return false;
}

bool detect_cycle_or_not(mi graph) {
    mb visited;
    
    for (const auto &p : graph) visited[p.first] = false;

    for (const auto &p : graph) {
        if (!visited[p.first]) {
            if (search(graph, visited, p.first, -1)) return true;
        }
    }
    return false;
}

int main() {
    mi graph;
    graph[1] = {2, 3};
    graph[2] = {1, 4};
    graph[3] = {1, 4};
    graph[4] = {2, 3};
    std::cout<<detect_cycle_or_not(graph)<<"\n";
}

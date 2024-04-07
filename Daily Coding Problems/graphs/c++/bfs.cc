#include <iostream>
#include <queue>
#include <list>
#include <string>
#include <map>

typedef std::string ss; 
typedef std::queue<ss> qs; 
typedef std::list<ss> ls; 
typedef std::map<ss, ls> msl;
typedef std::map<ss, bool> msb;

void bfs(msl graph, ss start, msb &visited) {
    qs q;

    q.push(start);
    while(!q.empty()) {
        ss head = q.front();
        q.pop();
        visited[head] = true;
        for (auto nei : graph[head]) {
            if (visited.find(nei) == visited.end())
                q.push(nei);
        }
    }
}


int main() {
    msl graph;
    msb visited;
    graph["JFK"] = {"SFO", "LAX"};
    graph["SFO"] = {"ORL"};
    graph["ORL"] = {"JFK", "LAX", "DFW"};
    graph["LAX"] = {"DFW"};

    bfs(graph, "ORL",  visited);
    for (auto p : visited) {
        std::cout<<p.first<<" "<<p.second<<"\n";
    }

}

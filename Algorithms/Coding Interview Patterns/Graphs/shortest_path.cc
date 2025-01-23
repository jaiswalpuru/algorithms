#include <iostream>
#include <queue>
#include <unordered_map>
#include <set>
#include <vector>

using namespace std;

vector<int> get_shortest_distance_from_node(vector<vector<int>> edges, int n, int start) {
    unordered_map<int, vector<vector<int>>> graph;
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int ,int>>> pq;
    vector<int> distance(n, INT_MAX);
    distance[start] = 0;
    pq.push({0, start});
    
    for (int i = 0; i < edges.size(); i++) {
        graph[edges[i][0]].push_back({edges[i][1], edges[i][2]});
        graph[edges[i][1]].push_back({edges[i][0], edges[i][2]});
    }

    while (!pq.empty()) {
        auto curr = pq.top();
        pq.pop();
        int dis = curr.first;
        int curr_node = curr.second;
        if (dis > distance[curr_node]) continue;
        for (auto nei : graph[curr_node]) {
            int nei_dis = dis + nei[1];
            if (nei_dis < distance[nei[0]]) {
                distance[nei[0]] = nei_dis;
                pq.push({nei_dis, nei[0]});
            }
        }
    }
    for (int i = 0; i < distance.size(); i++) {
        if (distance[i] == INT_MAX) distance[i] = -1;
    }
    return distance;
}


int main(int argc, char** argv) {
    int n = 6;
    vector<vector<int>> edges = {{0, 1, 5}, {0, 2, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 4}, {2, 4, 5}};
    int start = 0;
    vector<int> res = get_shortest_distance_from_node(edges, n, start);
    cout << "distance from " << start << "\n";
    for (int i = 0; i < n; i++) {
        cout << i << " : " << res[i] << "\n";
    }
    return 0;
}

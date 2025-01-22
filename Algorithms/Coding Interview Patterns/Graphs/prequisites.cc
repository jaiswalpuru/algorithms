#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>

using namespace std;

bool is_possible(int n, vector<vector<int>> pre_req) {
    vector<int> indegree(n, 0);
    unordered_map<int, vector<int>> graph;
    queue<int> q;

    for (int i = 0; i < pre_req.size(); i++) {
        indegree[pre_req[i][1]]++;
        graph[pre_req[i][0]].push_back(pre_req[i][1]);
    }

    int enrolled_course = 0;
    for (int i = 0; i < n; i++) {
        if (indegree[i] == 0) q.push(i);
    }

    while(!q.empty()) {
        int curr = q.front();
        q.pop();
        enrolled_course++;
        for (int i = 0; i < graph[curr].size(); i++) {
            indegree[graph[curr][i]]--;
            if (indegree[graph[curr][i]] == 0) q.push(graph[curr][i]);
        }
    }
    return enrolled_course == n;
}

int main(int argc, char** argv) {
    vector<vector<int>> pre_req = {{0, 1}, {1, 2}, {2, 1}};
    int n = 3;
    cout << is_possible(n, pre_req) << "\n";
    return 0;
}

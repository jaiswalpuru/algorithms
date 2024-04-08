#include <iostream>
#include <map>
#include <queue>
#include <vector>

typedef std::string st;
typedef std::vector<st> vs;
typedef std::queue<st> qs;
typedef std::map<st, vs> mss;
typedef std::map<st, int> msi;

vs topo_sort(mss &graph) {
    msi indeg;
    vs res;
    mss pre_req_to_course;
    qs q;

    for (const auto &p : graph) {
        indeg[p.first] += p.second.size();
        for (auto nei : graph[p.first]) {
            pre_req_to_course[nei].push_back(p.first);
        }
    }

    for (auto const &p : indeg) {
        if (p.second == 0) {
            q.push(p.first);
        }
    }
    
    while (!q.empty()) {
        st curr = q.front();
        q.pop();
        res.push_back(curr);

        for (auto nei : pre_req_to_course[curr]) {
            indeg[nei]--;
            if (indeg[nei] == 0) {
                q.push(nei);    
            }
        }
    }
    
    if (res.size() < pre_req_to_course.size()) return {};

    return res;
}


int main() {
    mss graph;
    graph["CSC300"] = {"CSC100", "CSC200"};
    graph["CSC200"] = {"CSC100"};
    graph["CSC100"] = {};

    vs res = topo_sort(graph);
    for (auto curr : res) std::cout<<curr<<"\n";
    return 0;
}

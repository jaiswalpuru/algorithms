#include <iostream>
#include <map>
#include <queue>
#include <vector>
#include <utility>
#include <string>
#include <set>

typedef std::string st;
typedef std::vector<st> vs;
typedef std::vector<std::pair<st,st>> vss;
typedef std::vector<std::pair<st, int>> vp;
typedef std::set<int> si;
typedef std::vector<int> vi;
typedef std::map<st, vi> mp;
typedef std::priority_queue<std::pair<double, std::pair<st, st>>, std::vector<std::pair<double, std::pair<st, st>>>, 
        std::greater<std::pair<double, std::pair<st, st>>>> pq;

int get_common_users_size(vi a, vi b) {
    si visited;
    int common = 0;

    for (int num : a) visited.insert(num);
    for (int num : b) {
        if (visited.find(num) != visited.end()) {
            common++;
        }
    }

    return common;
}

// vp is list of pairs with webiste, user
vss most_similar_websites(vp &graph, int k) {
    mp g;
    pq pri_q;
    vs sites;
    vss res;

    for (auto curr : graph) g[curr.first].push_back(curr.second);
    for (auto const &p : g) sites.push_back(p.first);
    
    for (int i = 0; i < sites.size() - 1 ; i++) {
        for (int j = i+1; j < sites.size(); j++) {
            double common = (double)get_common_users_size(g[sites[i]], g[sites[j]]);
            double score = common/(double)(g[sites[i]].size() + g[sites[j]].size() - (int)common);
            pri_q.push({score, {sites[i], sites[j]}});
            if (pri_q.size() > k) {
                pri_q.pop();
            }
        }
    }
    
    while (!pri_q.empty()) {
        res.push_back(pri_q.top().second);
        pri_q.pop();
    }

    return res;
}

int main() {
    vp p = {{"google", 1}, {"google", 3}, {"google", 5}, 
            {"pets", 1}, {"pets", 2},
            {"yahoo", 6}, {"yahoo", 2}, {"yahoo", 3}, {"yahoo", 4}, {"yahoo", 5},
            {"wiki", 4}, {"wiki", 5}, {"wiki", 6}, {"wiki", 7},
            {"bing", 1}, {"bing", 3}, {"bing", 5}, {"bing", 6}};
    vss res = most_similar_websites(p, 1);
    
    for (auto s : res) {
        std::cout<<s.first<< ", "<<s.second<<"\n";
    }
    return 0;
}

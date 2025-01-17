#include <iostream>
#include <vector>
#include <set>
#include <unordered_map>

using namespace std;

typedef struct GraphNode {
    int val;
    vector<struct GraphNode*> nei;
} Node;

void _helper(Node* g, set<int>& visited) {
    if (visited.find(g->val) != visited.end()) return;

    visited.insert(g->val);
    for (auto nei : g->nei) {
        _helper(nei, visited);
    }
}

void dfs_print(Node* g) {
    set<int> visited;
    _helper(g, visited);
    for (auto v : visited) cout << v << " ";
    cout << "\n";
}

Node* dfs(Node* g, unordered_map<Node*, Node*>& clone_map) {
    if (clone_map.find(g) != clone_map.end()) return clone_map[g];
    
    Node* clone_node = new Node;
    clone_node->val = g->val;
    clone_map[g] = clone_node;

    for (auto nei : g->nei) {
        Node* clone_nei = dfs(nei, clone_map);
        clone_node->nei.push_back(clone_nei);
    }
    return clone_node;
}

Node* clone_graph(Node* g) {
    if (g == nullptr) return nullptr;
    unordered_map<Node*, Node*> visited;
    return dfs(g, visited);
}

int main(int argc, char **argv) {
    Node* g1 = new Node;
    g1->val = 0;
    Node* g2 = new Node;
    g2->val = 1;
    Node* g3 = new Node;
    g3->val = 2;
    Node* g4 = new Node;
    g4->val = 3;
    g1->nei = {g2, g3};
    g2->nei = {g1, g3};
    g3->nei = {g1, g2, g4};
    g4->nei = {g3};
    Node* clone = clone_graph(g1);
    dfs_print(clone);
    return 0;
}

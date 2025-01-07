#include <iostream>
#include <set>
#include "linked_list.h"

using namespace std;

bool linked_list_loop_naive(Node* node) {
    if (node == nullptr) return false;

    set<Node*> visited;
    while (node) {
        if (visited.find(node) != visited.end()) return true;
        visited.insert(node);
        node = node->next;
    }
    return false;
}

bool linked_list_loop(Node* node) {
    if (node == nullptr) return false;

    Node* slow = node;
    Node* fast = node;

    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast) return true;
    }

    return false;
}

int main(int argc, char **argv) {
    Node* node = new Node;
    node->val = 0;
    node->next = new Node;
    node->next->val = 1;
    node->next->next = new Node;
    node->next->next->val = 2;
    node->next->next->next = new Node;
    node->next->next->next->val = 3;
    node->next->next->next->next = new Node;
    node->next->next->next->next->val = 4;
    node->next->next->next->next->next = new Node;
    node->next->next->next->next->next->val = 5;
    node->next->next->next->next->next->next = node->next->next->next;
    cout << linked_list_loop(node) << "\n";
    return 0;
}


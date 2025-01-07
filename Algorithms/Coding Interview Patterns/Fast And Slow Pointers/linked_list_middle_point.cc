#include <iostream>
#include "linked_list.h"

using namespace std;

Node* middle(Node* node) {
    Node* slow = node;
    Node* fast = node;

    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
    }
    return slow;
}

int main(int argc, char **argv) {
    Node* node = new Node;
    node->val = 1;
    node->next = new Node;
    node->next->val = 2;
    node->next->next = new Node;
    node->next->next->val = 4;
    node->next->next->next = new Node;
    node->next->next->next->val = 7;
    node->next->next->next->next = new Node;
    node->next->next->next->next->val = 3;
    cout << middle(node)->val << "\n";
    return 0;
}

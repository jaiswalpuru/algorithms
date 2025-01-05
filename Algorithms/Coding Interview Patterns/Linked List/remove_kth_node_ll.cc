#include <iostream>
#include "linked_list.h"

using namespace std;

void remove_kth_node(Node** head, int k) {
    Node* curr = *head;
    Node* sentinel = new Node;
    sentinel->val = -1;
    sentinel->next = curr;
    Node* trailer = sentinel;
    Node* leader = sentinel;
    
    while (k > 0) {
        leader = leader->next;
        k--;
        if (leader == NULL) return;
    }

    while (leader->next != NULL) {
        leader = leader->next;
        trailer = trailer->next;
    }
    
    trailer->next = trailer->next->next;
    *head = curr;
    return;
}


int main(int argc, char **argv) {
    Node* head = new Node;
    head->val = 1;
    head->next = new Node;
    head->next->val = 2;
    head->next->next = new Node;
    head->next->next->val = 4;
    head->next->next->next = new Node;
    head->next->next->next->val = 7;
    head->next->next->next->next = new Node;
    head->next->next->next->next->val = 3;

    remove_kth_node(&head, 2);
    while (head != NULL) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << "\n";
    return 0;
}

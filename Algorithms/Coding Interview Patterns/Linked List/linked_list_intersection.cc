#include <iostream>
#include "linked_list.h"

using namespace std;

Node* intersection(Node* h1, Node* h2) {
    
    Node* a = h1;
    Node* b = h2;
    while (a != b) {
        if (a != NULL) a = a->next;
        else a = h2;

        if (b != NULL) b = b->next;
        else b = h1;
    }
    return a;
}

int main(int argc, char **argv) {
    Node* head = new Node;
    Node* head2 = new Node;

    head->val = 1;
    head->next = new Node;
    head->next->val = 3;
    head->next->next = new Node;
    head->next->next->val = 4;
    head->next->next->next = new Node;
    head->next->next->next->val = 8;
    head->next->next->next->next = new Node;
    head->next->next->next->next->val = 7;
    head->next->next->next->next->next = new Node;
    head->next->next->next->next->next->val = 2;

    head2->val = 6;
    head2->next = new Node;
    head2->next->val = 4;
    head2->next->next = head->next->next->next;

    Node* res = intersection(head, head2);
    cout << res->val<<"\n";
    return 0;
}

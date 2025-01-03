#include <iostream>
#include "linked_list.h"

using namespace std;

void reverse_linked_list_iterative(Node** head) {
    if (head == NULL) return;

    Node* prev = NULL;;
    Node* curr = *head;

    while (curr != NULL) {
        Node *next_node = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next_node;
    }

    *head = prev;
}

Node* reverse_helper(Node *head) {
    if (head == nullptr || head->next == nullptr) return head;
    
    Node* curr = reverse_helper(head->next);
    head->next->next = head;
    head->next = NULL;
    return curr;
}

void reverse_linked_list_recursive(Node **head) {
    *head = reverse_helper(*head);
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
    reverse_linked_list_recursive(&head);

    while (head != NULL) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << "\n";
    return 0;
}

#include <iostream>

using namespace std;

typedef struct MultiNode {
    int val;
    struct MultiNode* next;
    struct MultiNode* child;
    MultiNode(int val) {
        this->val = val;
    }
} Node;

Node* flatten_list(Node *head) {
    if (!head) return head;

    Node* tail = head;
    while (tail->next) tail = tail->next;

    Node* curr = head;

    while (curr) {
        if (curr->child) {
            tail->next = curr->child;
            curr->child = nullptr;
            while (tail->next) tail = tail->next;
        }
        curr = curr->next;
    }

    return head;
}


int main(int argc, char** argv) {
    Node *head = new Node(1);
    head->next = new Node(2);
    head->next->next = new Node(3);
    head->next->next->next = new Node(4);
    head->next->next->next->next = new Node(5);
    head->next->child = new Node(6);
    head->next->child->next = new Node(7);
    head->next->child->next->child = new Node(10);
    head->next->next->next->child = new Node(8);
    head->next->next->next->child->next = new Node(9);
    head->next->next->next->child->child = new Node(11);
    Node* flat_list = flatten_list(head);

    while (flat_list) {
        cout << flat_list->val << " ";
        flat_list = flat_list->next;
    }
    cout << "\n";
    return 0;
}

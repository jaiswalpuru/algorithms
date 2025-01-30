#include<iostream>

using namespace std;

typedef struct LL {
    int val;
    struct LL* next;
    LL(int v) : val(v) {}
} Node;


Node* split_list(Node* head) {
    Node* slow = head;
    Node* fast = head;

    while (fast->next && fast->next->next) {
        slow = slow->next;
        fast = fast->next->next;
    }
    Node* second_head = slow->next;
    slow->next = nullptr;
    return second_head;
}

Node* merge(Node* a, Node* b) {
    Node* res = new Node(-1);
    Node* temp = res;

    while (a != nullptr && b != nullptr) {
        if (a->val > b->val) {
            temp->next = new Node(b->val);
            b = b->next;
        } else {
            temp->next = new Node(a->val);
            a = a->next;
        }
        temp = temp->next;
    }
    if (!a) temp->next = b;
    if (!b) temp->next = a;
    return res->next;
}

Node* sort_list(Node* head) {
    if (head == nullptr || head->next == nullptr) return head;

    Node* second_head = split_list(head);
    Node* first_part = sort_list(head);
    Node* sec_part = sort_list(second_head);
    return merge(first_part, sec_part);
}

int main(int argc, char** argv) {
    Node* head = new Node(3);
    head->next = new Node(2);
    head->next->next = new Node(4);
    head->next->next->next = new Node(5);
    head->next->next->next->next = new Node(1);
    Node* res = sort_list(head);
    while (res) {
        cout << res->val << " ";
        res = res->next;
    }
    cout << "\n";
    return 0;
}

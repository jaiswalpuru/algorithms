#include <iostream>
#include "linked_list.h"

using namespace std;

Node* find_mid(Node *head) {
    Node* fp = head;
    Node* sp = head;

    while (fp && fp->next) {
        sp = sp->next;
        fp = fp->next->next;
    }

    return sp;
}

Node* reverse(Node *head) {
    if (head == nullptr || head->next == nullptr) return head;

    Node* curr = reverse(head->next);
    head->next->next = head;
    head->next = nullptr;
    return curr;
}

bool is_palindrome(Node *head) {
    Node *mid = find_mid(head);
    Node *reverse_mid = reverse(mid);
    
    while (reverse_mid) {
        if (head->val != reverse_mid->val) return false;
        head = head->next;
        reverse_mid = reverse_mid->next;
    }
    return true;
}

int main(int argc, char** argv) {
    Node* head = new Node;
    head->val = 1;
    head->next = new Node;
    head->next->val = 2;
    head->next->next = new Node;
    head->next->next->val = 3;
    head->next->next->next = new Node;
    head->next->next->next->val = 2;
    head->next->next->next->next = new Node;
    head->next->next->next->next->val = 1;

    cout << is_palindrome(head) << "\n";
    return 0;
}


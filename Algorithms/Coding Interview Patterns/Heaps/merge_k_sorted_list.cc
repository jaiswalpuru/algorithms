#include <iostream>
#include <queue>

using namespace std;

typedef struct ListNode {
    int val;
    struct ListNode* next;
    ListNode(int v) : val(v) {} 
} Node;

Node* get_list(vector<Node*> arr) {
    auto cmp = [](const Node* a, const Node* b) {
        return a->val > b->val;
    };
    priority_queue<Node*, vector<Node*>, decltype(cmp)> pq(cmp);

    Node *res = new Node(-1);
    Node *head = res;
    for (int i = 0; i < arr.size(); i++) {
        pq.push(arr[i]);
    }

    while(!pq.empty()) {
        Node* curr = pq.top();
        pq.pop();
        head->next = curr;
        head = head->next;
        if (curr->next) pq.push(curr->next);
    }
        
    return res->next;;
}

int main(int argc, char** argv) {
    Node* a = new Node(1);
    a->next = new Node(6);
    Node* b = new Node(1);
    b->next = new Node(4);
    b->next->next = new Node(6);
    Node* c = new Node(3);
    c->next = new Node(7);
    vector<Node*> arr = {a, b, c};
    Node* ll = get_list(arr);
    while (ll) {
        cout << ll->val << " ";
        ll = ll->next;
    }
    cout << "\n";
    return 0;
}


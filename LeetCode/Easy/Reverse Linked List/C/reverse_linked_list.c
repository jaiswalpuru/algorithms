/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode ListNode;

void reverse(ListNode* node, ListNode** new_head) {
    if (node->next->next != NULL) reverse(node->next, new_head);
    else *new_head = node->next;
    node->next->next = node;
}

struct ListNode* reverseList(struct ListNode* head){
    ListNode* new_head = NULL;
    if (head == NULL || head->next == NULL) return head;
    reverse(head, &new_head);
    head->next = NULL;
    return new_head;
}

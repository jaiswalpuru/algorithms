/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode ListNode;

bool hasCycle(struct ListNode *head) {
    if (head == NULL || head->next == NULL) return false;
    ListNode *sp=head, *fp=head;
    while(fp->next != NULL && fp->next->next != NULL) {
        sp = sp->next;
        fp = fp->next->next;
        if (sp == fp) return true;
    }
    return false;
}

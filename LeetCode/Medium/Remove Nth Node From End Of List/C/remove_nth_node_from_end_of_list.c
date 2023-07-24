/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode ListNode;

ListNode* removeNthFromEnd(struct ListNode* head, int n){
    ListNode *fp = head;
    while(n > 0) {
        fp = fp->next;
        n--;
    }
    if (fp == NULL) return head->next;
    ListNode *sp=head;
    ListNode *prev = head;
    while(fp != NULL) {
        prev = sp;
        sp = sp->next;
        fp = fp->next;
    }
    ListNode *temp = sp;
    prev->next = sp->next;
    free(temp);
    return head;
}

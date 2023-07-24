/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode ListNode;

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    ListNode* res = (ListNode*) malloc(sizeof(ListNode));
    ListNode* prev = NULL;
    ListNode* head = (ListNode*) malloc(sizeof(ListNode));
    head->val = -1;
    head->next = res;
    int carry = 0;
    while(l1 != NULL && l2 != NULL) {
        int sum = l1->val+l2->val+carry;
        carry = sum/10;
        res->val = sum%10;
        prev = res;
        res->next = (ListNode*) malloc(sizeof(ListNode));
        res = res->next;
        l1 = l1->next;
        l2 = l2->next;
    }

    while(l1 != NULL) {
        int sum = l1->val+carry;
        carry = sum/10;
        res->val = sum%10;
        prev = res;
        res->next = (ListNode*) malloc(sizeof(ListNode));
        res = res->next;
        l1 = l1->next;
    }

    while(l2 != NULL) {
        int sum = l2->val+carry;
        carry = sum/10;
        res->val = sum%10;
        prev = res;
        res->next = (ListNode*) malloc(sizeof(ListNode));
        res = res->next;
        l2 = l2->next;
    }

    while(carry > 0) {
        res->val = carry%10;
        carry = carry/10;
        prev = res;
        res->next = (ListNode*) malloc(sizeof(ListNode));
        res = res->next;
    }
    prev->next = NULL;
    free(res);
    return head->next;
}

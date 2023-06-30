/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2){
    struct ListNode* res = (struct ListNode*) malloc(sizeof(struct ListNode));
    res->val = -1;
    res->next = NULL;
    struct ListNode* ans = res;
    while(list1 && list2) {
        res->next = (struct ListNode*) malloc(sizeof(struct ListNode));
        if (list1->val <= list2->val) {
            res->next->val = list1->val;
            list1 = list1->next;
        } else {
            res->next->val = list2->val;
            list2 = list2->next;
        }
        res = res->next;
        res->next = NULL;
    }
    while(list1) {
        res->next = (struct ListNode*) malloc(sizeof(struct ListNode));
        res->next->val = list1->val;
        res = res->next;
        res->next = NULL;
        list1 = list1->next;
    }
    while(list2) {
        res->next = (struct ListNode*) malloc(sizeof(struct ListNode));
        res->next->val = list2->val;
        res = res->next;
        res->next = NULL;
        list2 = list2->next;
    }
    return ans->next;
}

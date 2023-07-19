/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeNodes(struct ListNode* head){
  if (head == NULL || head->next == NULL) return head;
  struct ListNode* p = removeNodes(head->next);
  head->next = p;
  if (head->val < head->next->val) return head->next;
  return head;
}

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeNodes(ListNode* head) {
        if (!head) return nullptr;

        int sum = 0;
        ListNode *prev = head;
        ListNode *res = prev;
        ListNode *curr = head;

        while (curr) {
            if (curr->val == 0) {
                sum = curr->val;
                curr = curr->next;
                while (curr && curr->val != 0) {
                    sum += curr->val;
                    curr = curr->next;
                }
                prev->val = sum;
                prev = prev->next;
            } else {
                curr = curr->next;
            }
        }
        
        ListNode *temp = res;
        while (temp->next->next != prev) temp = temp->next;
        temp->next = nullptr;

        while (prev) {
            temp = prev;
            prev = prev->next;
            delete temp;
        }

        return res;
    }
};

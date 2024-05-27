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
    ListNode* removeNodes(ListNode* head) {
        ListNode* res;
        ListNode* temp;
        deque<ListNode*> dq;
        
        dq.push_back(head);
        head = head->next;
        
        while (head) {
            while (!dq.empty() && dq.back()->val < head->val) {
                dq.pop_back();
            }
            dq.push_back(head);
            head = head->next;
        }
        
        res = dq.front();
        temp = res;
        while (!dq.empty()) {
            temp->next = dq.front();
            temp = temp->next;
            dq.pop_front();
        }
        temp->next = nullptr;
        
        return res;
    }
};

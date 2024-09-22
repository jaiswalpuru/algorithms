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
    ListNode* insertGreatestCommonDivisors(ListNode* head) {
        ListNode* cur = head;
        
        while (cur->next != nullptr) {
            int gcd_val = gcd(cur->val, cur->next->val);
            ListNode *temp = cur->next;
            cur->next = new ListNode(gcd_val);
            cur->next->next = temp;
            cur = temp;
        }

        return head;
    }

    int gcd(int a, int b) {
        int t = 0;
        while (b > 0) {
            t = a;
            a = b;
            b = t % b;
        }
        return a;
    }
};

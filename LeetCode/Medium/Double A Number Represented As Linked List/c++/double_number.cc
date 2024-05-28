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
    ListNode* doubleIt(ListNode* head) {
        long long int carry = 0, val;

        double_linked_list(head, carry);
        if (carry > 0) {
            ListNode *temp = new ListNode(carry);
            temp->next = head;
            head = temp;
        }
        return head;
    }

    void double_linked_list(ListNode *head, long long int &carry) {
        if (head == nullptr) return;
        double_linked_list(head->next, carry);
        long long int cal = ((long long int)head->val * 2) + carry;
        head->val = (int)cal % 10;
        carry = (long long int)cal/10;
        return;
    }
};

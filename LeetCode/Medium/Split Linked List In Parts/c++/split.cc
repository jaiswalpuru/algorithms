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
    vector<ListNode*> splitListToParts(ListNode* head, int k) {
        int l = 0, b = 0, temp = 0;
        int parts = 0;
        ListNode* start = head;
        ListNode* curr = start;
        ListNode* prev = curr;
        vector<ListNode*> res;

        l = get_len(head);
        if (l > k) {
            b = l/k;
            parts = l % k;
        } else b = 1;

        temp = b;
        while (curr != nullptr) {
            if (temp == 0) {
                temp = b;
                if (parts-- > 0) {
                    prev = curr;
                    curr = curr->next;
                }
                prev->next = nullptr;
                res.push_back(start);
                start = curr;
            }
            prev = curr;
            curr = curr->next;
            temp--;
        }

        if (start != nullptr) res.push_back(start);
        while (res.size() < k) res.push_back(nullptr);

        return res;
    }

    int get_len(ListNode* head) {
        if (head == nullptr) return 0;
        return 1 + get_len(head->next);
    }
};

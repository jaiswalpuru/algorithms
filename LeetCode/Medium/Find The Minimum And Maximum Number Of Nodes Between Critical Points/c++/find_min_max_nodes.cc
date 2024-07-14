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
    vector<int> nodesBetweenCriticalPoints(ListNode* head) {
        vector<int> indexes, dis(2, -1);
        int i = 2;
        ListNode *prev = head;
        int min_dis = INT_MAX;

        head = head->next;
        if (head->next == nullptr) return dis;
        
        while (head->next != nullptr) {
            if (is_local_max(prev, head, head->next) || is_local_min(prev, head, head->next)) indexes.push_back(i);
            prev = head;
            head = head->next;
            i++;
        }

        if (indexes.size() > 0) {
            for (int j = 0; j < indexes.size()-1; j++) min_dis = min(min_dis, indexes[j+1]-indexes[j]);
        }
        
        dis[0] = min_dis == INT_MAX ? -1 : min_dis;
        dis[1] = indexes.size() > 1 ? indexes[indexes.size()-1] - indexes[0] : -1;
        
        return dis;
    }

    bool is_local_max(ListNode* prev, ListNode* curr, ListNode* next) { return prev->val < curr->val && curr->val > next->val; }

    bool is_local_min(ListNode* prev, ListNode* curr, ListNode* next) { return prev->val > curr->val && curr->val < next->val; }
};

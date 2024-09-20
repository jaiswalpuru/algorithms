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
    vector<vector<int>> spiralMatrix(int m, int n, ListNode* head) {
        vector<vector<int>> res(m, vector<int>(n));
        int left = 0;
        int right = n-1;
        int up = 0;
        int down = m-1;
        int t = 0;

        while (t < m * n) {
            for (int i = left; i <= right; i++) {
                if (head == nullptr) {
                    res[up][i] = -1;
                } else {
                    res[up][i] = head->val;
                    head = head->next;
                }
                t++;
            }

            for (int i = up+1; i <= down; i++) {
                if (head == nullptr) {
                    res[i][right] = -1;
                } else {
                    res[i][right] = head->val;
                    head = head->next;
                }
                t++;
            }

            if (up != down) {
                for (int i = right - 1; i >= left; i--) {
                    if (head == nullptr) {
                        res[down][i] = -1;
                    } else {
                        res[down][i] = head->val;
                        head = head->next;
                    }
                    t++;
                }
            }
            
            if (left != right) {
                for (int i = down-1; i > up; i--) {
                    if (head == nullptr) {
                        res[i][left] = -1;
                    } else {
                        res[i][left] = head->val;
                        head = head->next;
                    }
                    t++;
                }
            }

            left++;
            right--;
            up++;
            down--;
        }
        return res;
    }
};

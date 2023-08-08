/**
 * Definition for a category handler.
 * class CategoryHandler {
 * public:
 *     CategoryHandler(vector<int> categories);
 *     bool haveSameCategory(int a, int b);
 * };
 */
class Solution {
public:
    int _find(int x, vector<int>& parent) {
        if (x != parent[x]) parent[x] = _find(parent[x], parent);
        return parent[x];
    }

    void _union(int x, int y, vector<int>& parent, vector<int>& size, int& cnt) {
        x = _find(x, parent);
        y = _find(y, parent);
        if (x == y) return;
        if (size[x] > size[y]) {
            size[x] += size[y];
            parent[y] = x;
        } else {
            size[y] += size[x];
            parent[x] = y;
        }
        cnt--;
    }

    int numberOfCategories(int n, CategoryHandler* categoryHandler) {
        vector<int> parent(n);
        vector<int> size(n);
        int cnt = n;
        for (int i=0; i<n; i++) parent[i] = i;
        for (int i=0; i<n; i++) size[i] = 1;
        for (int i=0; i<n; i++) {
            for (int j=i+1; j<n; j++) {
                if (categoryHandler->haveSameCategory(i, j)) {
                    _union(i, j, parent, size, cnt);
                }
            }
        }
        return cnt;
    }
};

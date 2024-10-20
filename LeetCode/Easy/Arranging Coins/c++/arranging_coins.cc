class Solution {
public:
    int arrangeCoins(int n) {
        int cmp = 0;
        int tmp = n;

        for (int i = 1; i <= n; i++) {
            tmp -= i;
            if (tmp < 0) break;
            cmp++;
        }

        return cmp;
    }
};

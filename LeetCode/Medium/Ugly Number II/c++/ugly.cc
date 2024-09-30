class Solution {
public:
    int nthUglyNumber(int n) {
        if (n <= 6) return n;
        vector<int> res(n);
        int ind_2 = 0;
        int ind_3 = 0;
        int ind_5 = 0;

        res[0] = 1;
        for (int i = 1; i < n; i++) {
            res[i] = min(res[ind_2] * 2, min(res[ind_3] * 3, res[ind_5] * 5));
            if ((res[ind_2] * 2) == res[i]) ind_2++;
            if ((res[ind_3] * 3) == res[i]) ind_3++;
            if ((res[ind_5] * 5) == res[i]) ind_5++;
        }

        return res[res.size()-1];
    }
};

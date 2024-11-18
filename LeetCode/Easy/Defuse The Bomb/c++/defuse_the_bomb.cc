class Solution {
public:
    vector<int> decrypt(vector<int>& code, int k) {
        int n = code.size();
        vector<int> res(code.size());

        for (int i = 0; i < n; i++) {
            if (k == 0) {
                res[i] = 0;
            } else if (k > 0) {
                int j = i + 1;
                int t = k;
                while (t > 0) {
                    res[i] += code[j % n];
                    j++;
                    t--;
                }
            } else {
                int j = i - 1;
                int t = k;
                while (t < 0) {
                    if (j < 0) {
                        j = n - 1;
                    }
                    res[i] += code[j];
                    j--;
                    t++;
                }
            }
        }

        return res;
    }
};

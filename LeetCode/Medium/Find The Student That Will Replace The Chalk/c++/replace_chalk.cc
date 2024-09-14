class Solution {
public:
    int chalkReplacer(vector<int>& chalk, int k) {
        long long sum_chalk = 0;
        for (int i = 0; i < chalk.size(); i++) sum_chalk += chalk[i];

        int i = 0;
        while (k >= sum_chalk) k -= sum_chalk;
        for (i = 0; i < chalk.size(); i++) {
            if (k >= chalk[i]) k -= chalk[i];
            else break;
        }

        return i;
    }
};

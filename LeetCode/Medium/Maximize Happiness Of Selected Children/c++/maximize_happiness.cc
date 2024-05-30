class Solution {
public:
    long long maximumHappinessSum(vector<int>& happiness, int k) {
        
        int i = 0;
        long long happiness_sum = 0;
        
        sort(happiness.begin(), happiness.end(), greater<int>());
        while (k > 0) {
            if (happiness[i] > i) {
                happiness_sum += happiness[i] - i;
            }
            i++;
            k--;
        }

        return happiness_sum;
    }
};

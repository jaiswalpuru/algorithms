class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        int i=0, j=0;
        double sum = 0;
        double avg = -100000000;
        while(j < nums.size()) {
            if (j-i > k) sum -= nums[i++];
            if (j-i == k) avg = max(avg, sum/k);
            sum += nums[j];
            j++;
        }
        if (j-i > k) sum -= nums[i++];
        if (j-i == k) avg = max(avg, sum/k);
        return avg;
    }
};

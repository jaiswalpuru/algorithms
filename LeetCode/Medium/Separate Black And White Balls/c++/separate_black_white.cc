class Solution {
public:
    long long minimumSteps(string s) {
        long long min_steps = 0LL;
        int num_zeros = 0;
        
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s[i] == '0') {
                num_zeros++;
            } else {
                min_steps += num_zeros;
            }
        }

        return min_steps;
    }
};

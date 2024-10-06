class Solution {
public:
    int countNicePairs(vector<int>& nums) {
        int cnt = 0;
        long long temp = 0LL;
        int MOD = (int)1e9 + 7;
        unordered_map<int, int> num_map;

        for (int num : nums) num_map[num - rev(num)] = (num_map[num - rev(num)] + 1) % MOD;
        for (auto p : num_map) {
            temp = ((long long)p.second) * ((long long)p.second - 1) / 2;
            if (p.second > 1) cnt += temp % MOD;
        }
        
        return cnt % MOD;
    }

    int rev(int num) {
        int n = 0;
        while (num > 0) {
            n = (n * 10) + (num % 10);
            num /= 10;
        }
        return n;
    }
};

class Solution {
public:
    bool canArrange(vector<int>& arr, int k) {
        unordered_map<int, int> rem_cnt;

        for (int num : arr) rem_cnt[(num % k + k) % k]++;

        for (int num : arr) {
            int rem = (num % k + k) % k;
            if (rem == 0) {
                if (rem_cnt[rem] % 2) return false;
            }
            else if (rem_cnt[rem] != rem_cnt[k - rem]) return false; 
        }

        return true;
    }
};

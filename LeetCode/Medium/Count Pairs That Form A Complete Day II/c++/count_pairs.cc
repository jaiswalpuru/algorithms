class Solution {
public:
    long long countCompleteDayPairs(vector<int>& hours) {
        long long res = 0;
        unordered_map<int, int> cnt;

        for (int i = 0; i < hours.size(); i++) {
            res += cnt[(24 - hours[i] % 24) % 24];
            cnt[hours[i] % 24]++;
        }

        return res;
    }
};

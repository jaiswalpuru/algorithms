class Solution {
public:
    long long dividePlayers(vector<int>& skill) {
        vector<pair<int, int>> pairs(skill.size()/2);
        sort(skill.begin(), skill.end());
        int l = 0, r = skill.size()-1;
        long long sum = 0LL;
        while (l < r) {
            if (sum == 0) {
                sum = skill[l] + skill[r];
                pairs.push_back(make_pair(skill[l], skill[r]));
            } else {
                if (sum != (skill[l] + skill[r])) return -1;
                pairs.push_back(make_pair(skill[l], skill[r]));
            }
            l++; r--;
        }
        
        sum = 0;
        for (pair<int, int> p : pairs) {
            sum += (p.first * p.second);
        }

        return sum;
    }
};

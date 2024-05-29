class Solution {
public:
    vector<string> findRelativeRanks(vector<int>& score) {
        vector<int> score_copy(score);
        vector<string> vs;
        map<int, string> mapping;

        sort(score_copy.begin(), score_copy.end(), std::greater<int>());
        
        for (int i = 0; i < score_copy.size(); i++) {
            if (i  == 0) mapping[score_copy[i]] = "Gold Medal";
            else if (i == 1) mapping[score_copy[i]] = "Silver Medal";
            else if (i == 2) mapping[score_copy[i]] = "Bronze Medal";
            else mapping[score_copy[i]] = std::to_string(i+1);
        }

        for (int i = 0; i < score.size(); i++) {
           vs.push_back(mapping[score[i]]);
        }

        return vs;
    }
};

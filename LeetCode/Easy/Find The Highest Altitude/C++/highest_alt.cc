class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        int sum = 0;
        int curr = gain[0];
        sum = max(sum, curr);
        for (int i=1; i<gain.size(); i++) {
            curr += gain[i];
            sum = max(sum, curr);
        }
        return sum;
    }
};

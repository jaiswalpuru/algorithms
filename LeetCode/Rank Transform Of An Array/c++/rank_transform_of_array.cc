class Solution {
public:
    vector<int> arrayRankTransform(vector<int>& arr) {
        int rank = 1;
        vector<int> temp = arr;
        unordered_map<int, int> val_rank;
        sort(temp.begin(), temp.end());
        for (int i = 0; i < arr.size(); i++) {
            if (val_rank[temp[i]] == 0) {
                val_rank[temp[i]] = rank;
                rank++;
            }
        }

        for (int i = 0; i < arr.size(); i++) temp[i] = val_rank[arr[i]];

        return temp;
    }
};

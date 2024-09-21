class Solution {
public:
    vector<int> xorQueries(vector<int>& arr, vector<vector<int>>& queries) {
        int j = 0;
        vector<int> res;
        vector<int> prefix(arr.size()+1);

        prefix[0] = 0;
        for (int i = 0; i < arr.size(); i++) {
            prefix[j+1] = (prefix[j] ^ arr[i]);
            j++;
        }

        for (int i = 0; i < queries.size(); i++) {
            int l = queries[i][0];
            int r = queries[i][1];
            res.push_back(prefix[l] ^ prefix[r+1]);
        }
        return res;
    }
};

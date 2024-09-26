class Solution {
public:
    vector<int> lexicalOrder(int n) {
        vector<int> res;
        for (int i = 1; i <= 9; i++) {
            generate(i, n, res);
        }
        return res;
    }

    void generate(int i, int limit, vector<int> &res) {
        if (i > limit) return;

        res.push_back(i);

        for (int j = 0; j <= 9; j++) {
            int next_num = i * 10 + j;
            if (next_num <= limit) generate(next_num, limit, res);
            else break;
        }
    }
};

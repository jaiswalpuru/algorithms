class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        generate_paran(n, res, 0, 0, "");
        return res;
    }

    void generate_paran(int n, vector<string>& res, int i, int j, string s) {
        if (i + j == 2 * n) {
            res.push_back(s);
            return;
        }

        if (i < n) {
            s += "(";
            generate_paran(n, res, i + 1, j, s);
            s.pop_back();
        }
        
        if (j < i) {
            s += ")";
            generate_paran(n, res, i, j + 1, s);
            s.pop_back();
        }
    }
};

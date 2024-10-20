class Solution {
public:
    bool parseBoolExpr(string expression) {
        int ind = 0;
        return eval(expression, ind);
    }

    bool eval(string& expr, int& i) {
        char c = expr[i++];

        if (c == 't') return true;
        if (c == 'f') return false;

        // for ! op
        if (c == '!') {
            i++; // skip (
            bool res = !eval(expr, i);
            i++; // skip )
            return res;
        }

        // &, |
        vector<bool> val;
        i++;
        while (expr[i] != ')') {
            if (expr[i] != ',') {
                val.push_back(eval(expr, i));
            } else {
                i++;
            }
        }
        i++; // skip )

        if (c == '&') {
            for (bool v : val) {
                if (!v) return false;
            }
            return true;
        }

        if (c == '|') {
            for (bool v : val) {
                if (v) return true;
            }
            return false;
        }

        return false;
    }
};

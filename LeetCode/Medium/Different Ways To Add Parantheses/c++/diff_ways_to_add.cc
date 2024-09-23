class Solution {
public:
    vector<int> diffWaysToCompute(string expression) {
        map<string, vector<int>> cache;
        return diffWaysToComputeMemo(expression, cache);
    }

    vector<int> diffWaysToComputeMemo(string exp, map<string, vector<int>> &cache) {
        if (cache.find(exp) != cache.end()) return cache[exp];
        vector<int> res;

        for (int i = 0; i < exp.length(); i++) {
            if (exp[i] == '-' || exp[i] == '+' || exp[i] == '*') {
                string left = exp.substr(0, i);
                string right = exp.substr(i+1, exp.length());
                vector<int> left_exp = diffWaysToComputeMemo(left, cache);
                vector<int> right_exp = diffWaysToComputeMemo(right, cache);
                for (int j = 0; j < left_exp.size(); j++) {
                    for (int k = 0; k < right_exp.size(); k++) {
                        int val = 0;
                        switch (exp[i]) {
                            case '+':
                                val = left_exp[j] + right_exp[k];
                                break;
                            case '-':
                                val = left_exp[j] - right_exp[k];
                                break;
                            case '*':
                                val = left_exp[j] * right_exp[k];
                                break;
                        }
                        res.push_back(val);
                    }
                }
            }
        }
        
        if (res.size() == 0) res.push_back(stoi(exp));
        cache[exp] = res;
        return res;
    }
};

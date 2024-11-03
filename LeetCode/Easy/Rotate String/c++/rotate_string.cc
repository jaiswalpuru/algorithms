class Solution {
public:
    bool rotateString(string s, string goal) {
        if (s == goal) return true;
        if (s.length() != goal.length()) return false;

        for (int i = 0; i < s.length(); i++) {
            if (s.substr(0, i) == goal.substr(goal.length() - i, goal.length()) &&
                s.substr(i, s.length()) == goal.substr(0, goal.length() - i)) return true;
        }

        return false;
    }
};

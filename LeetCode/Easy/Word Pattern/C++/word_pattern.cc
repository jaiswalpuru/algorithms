class Solution {
public:
    bool wordPattern(string pattern, string s) {
        map<char, string> char_word;
        map<string, char> word_char;
        vector<string> vs;
        stringstream ss(s);
        string word;
        while(ss >> word) vs.push_back(word);
        if (vs.size() != pattern.length()) return false;
        for (int i=0; i<vs.size(); i++) {
            if (char_word.find(pattern[i]) != char_word.end()) {
                if (char_word.find(pattern[i])->second != vs[i]) return false;
            } else {
                char_word[pattern[i]] = vs[i];
            }
            if (word_char.find(vs[i]) != word_char.end()) {
                if (word_char.find(vs[i])->second != pattern[i]) return false;
            } else {
                word_char[vs[i]] = pattern[i];
            }
        }
        return true;
    }
};

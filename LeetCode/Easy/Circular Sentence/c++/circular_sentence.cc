class Solution {
public:
    bool isCircularSentence(string sentence) {
        if (sentence[0] != sentence[sentence.length()-1]) return false;
        vector<string> st;
        string word;
        stringstream ss(sentence);

        while (ss >> word) {
            st.push_back(word);
        }

        for (int i = 0; i < st.size() - 1; i++) {
            if (st[i][st[i].length()-1] != st[i + 1][0]) return false;
        }

        return true;
    }
};

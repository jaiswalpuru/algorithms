class Solution {
public:
    string mergeAlternately(string word1, string word2) {
        string merged = "";
        if (word1.length() == 0) return word2;
        if (word2.length() == 0) return word1;
        int i=0, j=0;
        while(i < word1.length() && j < word2.length()) {
            merged += word1[i++];
            merged += word2[j++];
        }
        if (i < word1.length()) merged += word1.substr(i, word1.length());
        if (j < word2.length()) merged += word2.substr(j, word2.length());
        return merged;
    }
};

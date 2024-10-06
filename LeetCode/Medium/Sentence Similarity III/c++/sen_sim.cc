class Solution {
public:
    bool areSentencesSimilar(string sentence1, string sentence2) {
        vector<string> sen1, sen2;
        stringstream s1(sentence1);
        stringstream s2(sentence2);
        string word;

        while (s1 >> word) sen1.push_back(word);
        while (s2 >> word) sen2.push_back(word);

        while (!sen1.empty() && !sen2.empty()) {
            if (sen1.front() == sen2.front()) {
                sen1.erase(sen1.begin());
                sen2.erase(sen2.begin());
            } else if (sen1.back() == sen2.back()){
                sen1.pop_back();
                sen2.pop_back();              
            } else {
                break;
            }
        }

        return sen1.empty() || sen2.empty();
    }
};

class Solution {
public:
    int compareVersion(string version1, string version2) {
        stringstream ss1(version1);
        stringstream ss2(version2);
        vector<int> v1, v2;
        string word;
        bool is_true;
        int i = 0, j = 0;
        int num = 0;

        while(!ss1.eof()) {
            getline(ss1, word, '.');
            num = stoi(word);
            v1.push_back(num);
        }

        while(!ss2.eof()) {
            getline(ss2, word, '.');
            num = stoi(word);
            v2.push_back(num);
        }

        while(i < v1.size() && j < v2.size()) {
            if (v1[i] > v2[j]) return 1;
            else if (v1[i] < v2[j]) return -1;
            i++; j++;
        }
    
        if (i == v1.size() && j == v2.size()) return 0;
        else if (i == v1.size()) {
            is_true = true;
            while (j < v2.size()) {
                if (v2[j] != 0) {
                    is_true = false;
                    break;
                }
                j++;
            }
            return is_true ? 0 : -1;
        } else {
            is_true = true;
            while (i < v1.size()) {
                if (v1[i] != 0) {
                    is_true = false;
                    break;
                }
                i++;
            }
            return is_true ? 0 : 1;
        }
    }
};

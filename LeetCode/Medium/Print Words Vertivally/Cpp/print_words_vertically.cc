#include <iostream>
#include <vector>
#include <sstream>

#define PB push_back 

typedef std::vector<std::string> vs;

vs split(std::string s, char delim, int &len) {
    vs st;
    std::stringstream ss(s);
    std::string token;
    while(getline(ss, token, delim)) {
        len = std::max(len, (int)token.size());
        st.PB(token);
    }
    return st;
}

vs print_words_vertically(std::string s) {
    int len = 0;
    vs split_words = split(s, ' ', len);
    vs res;
    for (int i=0;i<len;i++) {
        std::string temp = "";
        for (int j=0; j<split_words.size(); j++) {
            if (split_words[j].size() <= i) {
                temp += " ";
                continue;
            }
            temp += split_words[j][i];
        }
        int k = 0;
        for (k=temp.size()-1; k>=0; k--) {
            if (temp[k] != ' ') break;
        }
        std::string t = "";
        for (int j=0;j<=k;j++) {
            t += temp[j];
        }
        res.PB(temp);
    }
    return res;
}

int main() {
    vs res = print_words_vertically("HOW ARE YOU");
    for (auto it=res.begin(); it!= res.end(); it++) std::cout<<*it<<"\n";
}

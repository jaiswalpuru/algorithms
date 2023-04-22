#include <iostream>
#include <map>
#include <vector>

typedef std::vector<std::string> vs;
typedef std::map<std::string, int> msi;

int longest_palindrome_by_concatenating_two_letter_words(vs &v) {
    msi cnt;
    std::string s = "";
    int odd = 0, ans = 0;
    for (int i=0;i<v.size(); i++) cnt[v[i]]++;
    for (auto it=cnt.begin(); it!=cnt.end(); it++){
        std::string word = it->first;
        if (word[0] == word[1]) {
            if (cnt[word]%2 && cnt[word] > odd) {
                s = word;
                odd = cnt[word];
            }
        }
    }
    ans += odd*2;
    cnt[s] = 0;
    //handle words with same characters
    for (int i=0;i<v.size(); i++) {
        std::string word = v[i];
        int c = cnt[word];
        if (word[0] == word[1] && c > 0) {
            if (c % 2)  ans += (c-1)*2;
            else ans += c * 2;
            cnt[word] = 0;
        }
    }
    //handle words with different characters
    for (int i=0;i<v.size();i++){
        std::string word = v[i];
        int c = cnt[word];
        if (c > 0) {
            std::string rev = word;
            reverse(rev.begin(), rev.end());
            ans += std::min(cnt[word], cnt[rev])*4;
            cnt[word] = 0;
            cnt[rev] = 0;
        }
    }
    return ans;
}

int main() {
    vs v{"nn","nn","hg","gn","nn","hh","gh","nn","nh","nh"};
    std::cout<<longest_palindrome_by_concatenating_two_letter_words(v)<<"\n";
} 

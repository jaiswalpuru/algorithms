#include <iostream>
#include <string>
#include <map>

typedef std::string st;
typedef std::map<char, int> mci;

int longest_substring_without_repeating_chars(st s) {
    mci cnt;
    int i = 0, j = 0;
    int gbl_max = 0;

    while (i < s.size()) {
        cnt[s[i]]++;

        if (cnt[s[i]] > 1) {
            gbl_max = std::max(gbl_max, i-j);
            
            while (j <= i && cnt[s[i]] > 1) {
                cnt[s[j]]--;
                j++;
            }
        }
        i++;
    }
    
    gbl_max = std::max(gbl_max, i-j);
    return gbl_max;
}

int main() {
    
    st s = "abcabcbb";
    std::cout<<longest_substring_without_repeating_chars(s) << "\n";

    s = "bbbbb";
    std::cout<<longest_substring_without_repeating_chars(s) << "\n";

    s = "pwwkew";
    std::cout<<longest_substring_without_repeating_chars(s) << "\n";

    s = "";
    std::cout<<longest_substring_without_repeating_chars(s) << "\n";

    s = "abba";
    std::cout<<longest_substring_without_repeating_chars(s) << "\n";

    return 0;
}

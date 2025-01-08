#include <iostream>

using namespace std;

int substring_anagrams(string s, string t) {
    int f_s[26] = {0};
    int f_t[26] = {0};
    int l = 0;
    int r = 0;
    int res = 0;

    for (int i = 0; i < t.length(); i++) f_t[t[i] - 'a']++;

    while (r < s.length()) {
        f_s[s[r] - 'a']++;
        if (r - l + 1 == t.length()) {
            bool is_same = true;
            for (int i = 0; i < 26; i++) {
                if (f_s[i] != f_t[i]) {
                    is_same = false;
                    break;
                }
            }
            if (is_same) res++;
            f_s[s[l] - 'a']--;
            l++;
        }
        r++;
    }
    return res;
}
    

int main(int argc, char **argv) {
    cout << substring_anagrams("caabab", "aba") << "\n";
    return 0;
}

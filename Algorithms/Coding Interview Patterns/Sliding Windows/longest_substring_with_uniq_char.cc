#include <iostream>
#include <set>
#include <unordered_map>

using namespace std;

int longest_substring(string s) {
    int res = 0;
    set<char> visited;
    int l = 0;
    int r = 0;

    while (r < s.length()) {
        while (visited.find(s[r]) != visited.end()) {
            visited.erase(s[l]);
            l++;
        }

        res = max(res, r - l + 1);
        visited.insert(s[r]);
        r++;
    }

    return res;
}

int longest_substring_optimize(string s) {
    unordered_map<char, int> prev_ind;
    int l = 0;
    int r = 0;
    int res = 0;

    while (r < s.length()) {
        if (prev_ind.find(s[r]) != prev_ind.end() && prev_ind[s[r]] >= l) {
            l = prev_ind[s[r]] + 1;
        }
        res = max(res, r - l + 1);
        prev_ind[s[r]] = r;
        r++;
    }
    return res;
}

int main(int argc, char **argv) {
    cout << longest_substring_optimize("abcba") << "\n";
    return 0;
}

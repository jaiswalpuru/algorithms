#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

int longest_substring(string s, int k) {
    unordered_map<char, int> freq;
    int high_freq = 0;
    int max_len = 0;
    int left = 0;
    int right = 0;
    int num_chars_to_replace = 0;

    while (right < s.length()) {
        freq[s[right]]++;
        high_freq = max(high_freq, freq[s[right]]);
        num_chars_to_replace = (right - left + 1) - high_freq;

        if (num_chars_to_replace > k) {
            freq[s[left]]--;
            left++;
        }

        max_len = right - left + 1;
        right++;
    }
    return max_len;
}

int main(int argc, char **argv) {
    cout << longest_substring("aabcdcca", 2) << "\n";
    return 0;
}

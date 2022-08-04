#include <iostream>
#include <string>
#include <map>
#include <algorithm>

using namespace std;

typedef long long ll;
typedef map<char, ll> mll;

void palindrome_reorder(string s) {
    mll cnt;
    int odd_cnt=0;
    for (int i=0;i<s.length(); i++){
        cnt[s[i]]++;
    }
    for (auto it: cnt){
        if (it.second%2){
            odd_cnt++;
        }
    }
    if (odd_cnt > 1) {
        cout<<"NO SOLUTION\n";
    } else {
        string s = "", odd_char = "";
        for (auto it: cnt) {
            if (it.second%2 == 0){
                for (int i=0;i<it.second/2;i++){
                    s += it.first;
                }
            }else {
                for (int i=0;i<it.second;i++){
                    odd_char += it.first;
                }
            }
        }
        s += odd_char;
        string rev_string = s;
        if (odd_char!=""){
            rev_string = rev_string.substr(0, s.length()-cnt[odd_char[0]]);
        }
        reverse(rev_string.begin(), rev_string.end());
        s += rev_string;
        cout<<s<<"\n";
    }
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    string s;
    cin>>s;
    palindrome_reorder(s);
}

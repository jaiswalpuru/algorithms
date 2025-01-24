#include <iostream>
#include <unordered_map>

using namespace std;

void dfs(int ind, string st, string key, unordered_map<char, string> phone, vector<string>& res) {
    if (st.length() == key.length()) {
        res.push_back(st);
        return;
    }

    for (auto k : phone[key[ind]]) {
        st += k;
        dfs(ind + 1, st, key, phone, res);
        st = st.substr(0, st.length()-1);
    }
}

vector<string> phone_key(string key, unordered_map<char, string> phone) {
    vector<string> res;
    dfs(0, "", key, phone, res);
    return res;
}

int main(int argc, char** argv) {
    unordered_map<char, string> phone = {{'2', "abc"}, {'3', "def"}, {'4', "ghi"}, {'5', "jkl"}, {'6', "mno"}, {'7', "pqrs"}, {'8', "tuv"}, {'9', "wxyz"}};
    vector<string> res = phone_key("69", phone);
    for (auto s : res) cout << s << "\n";
    return 0;
}

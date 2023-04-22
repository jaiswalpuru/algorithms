#include <iostream>
#include <stack>

typedef std::stack<char> sc;

std::string remove_all_adjacent_duplicates_in_string(std::string s) {
    sc st;
    for (int i=0;i<s.size();i++) {
        if (st.empty()) {
            st.push(s[i]);
            continue;
        }
        if (!st.empty() && st.top() == s[i]) {
            while(!st.empty() && st.top() == s[i]) st.pop();
        }else {
            st.push(s[i]);
        }
    }
    std::string res = "";
    while(!st.empty()) {
        res = st.top()+res;
        st.pop();
    }
    return res;
}

int main() {
    std::cout<<remove_all_adjacent_duplicates_in_string("abbaca")<<"\n";
}

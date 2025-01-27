#include <iostream>
#include <stack>

using namespace std;

vector<int> next_largest_element(vector<int> arr) {
    stack<int> st;
    vector<int> res(arr.size());

    for (int i = arr.size() - 1; i >= 0; i--) {
        while (!st.empty() && st.top() <= arr[i]) st.pop();
        res[i] = !st.empty() ? st.top() : -1;
        st.push(arr[i]);
    }
    return res;
}

int main(int argc, char** argv) {
    vector<int> res = next_largest_element({5, 2, 4, 6, 1});
    for (int i = 0; i < res.size(); i++) cout << res[i] << " ";
    cout << "\n";
    return 0;
}

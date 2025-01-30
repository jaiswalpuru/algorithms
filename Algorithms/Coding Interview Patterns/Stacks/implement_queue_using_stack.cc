#include <iostream>
#include <stack>

using namespace std;

void transfer_enq_to_deq(stack<int>& enq_st, stack<int>& deq_st) {
    if (deq_st.empty()) {
        while (!enq_st.empty()) {
            deq_st.push(enq_st.top());
            enq_st.pop();
        }
    }
}

void enq(stack<int>& enq_st, int val) {
    enq_st.push(val);
}

int deq(stack<int>& enq_st, stack<int>& deq_st) {
    transfer_enq_to_deq(enq_st, deq_st);
    if (deq_st.empty()) return -1;
    int val =  deq_st.top();
    deq_st.pop();
    return val;
}

int peek(stack<int>& enq_st, stack<int>& deq_st) {
    transfer_enq_to_deq(enq_st, deq_st);
    if (deq_st.empty()) return -1;
    return deq_st.top();
}

void perform_queries(vector<vector<string>> q) {
    stack<int> enq_st;
    stack<int> deq_st;

    for (int i = 0; i < q.size(); i++) {
        if (q[i][0] == "enq") {
            enq(enq_st, stoi(q[i][1]));
        } else if (q[i][0] == "deq") {
            cout << "deq " << deq(enq_st, deq_st) << "\n";
        } else if (q[i][0] == "peek") {
            cout << "peek " << peek(enq_st, deq_st) << "\n";
        }
    }
}

int main(int argc, char** argv) {
    vector<vector<string>> queries = {{"enq", "1"}, {"enq", "2"}, {"deq", "-1"}, {"enq", "3"}, {"peek", "-1"}};
    perform_queries(queries);
    return 0;
}

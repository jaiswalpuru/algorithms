#include <iostream>
#include <cstdlib>
#include <algorithm>
#include <queue>
#include <string>
#include <map>
using namespace std;

string repeat_limited_string(string s, int repeat_limit) {
    string res = "";
    map<char,int> count;

    for (auto it : s){
        count[it]++;
    }

    priority_queue<pair<char,int> > pq;

    for (auto it : count){
        pq.push(make_pair(it.first,it.second));
    }

    while(!pq.empty()){
        char c1 = pq.top().first;
        int cnt_c1 = pq.top().second;
        pq.pop();

        int len = min(repeat_limit, cnt_c1);
        for (int i=0;i<len;i++){
            res+=c1;
        }

        if (!pq.empty() and cnt_c1-len>0){
            char c2 = pq.top().first;
            int cnt_c2 = pq.top().second;
            pq.pop();

            res+=c2;
            pq.push(make_pair(c1,cnt_c1-len));
            if (cnt_c2-1>0){
                pq.push(make_pair(c2, cnt_c2-1));
            }
        }
    }
    return res;
}


int main(){
    cout<<repeat_limited_string("cczazcc", 3)<<"\n";
    return 0;
}
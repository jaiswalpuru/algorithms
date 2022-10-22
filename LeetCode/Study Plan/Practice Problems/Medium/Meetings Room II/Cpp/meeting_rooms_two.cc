#include <iostream>
#include <vector>
#include <queue>

typedef std::vector<int> vi;
typedef std::vector<vi> vvi;
typedef std::priority_queue<int, vi, std::greater<int>> pqi;

int meeting_rooms_two(vvi &v) {
    std::sort(v.begin(), v.end(), [](const vi& v1, const vi& v2){
        return v1[0] < v2[0];
    });
    int rooms = 1;
    pqi p;
    p.push(v[0][1]);
    for (int i=1; i<v.size(); i++){
        int top = p.top();
        if (v[i][0] < top) {
            rooms++;
        }else {
            p.pop();
        }
        p.push(v[i][1]);
    }
    return rooms;
}

int main() {
    vvi v
    {
        {0,30},
        {5,10},
        {15, 20}
    };
    std::cout<<meeting_rooms_two(v)<<"\n";
}

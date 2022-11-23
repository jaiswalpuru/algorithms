#include <iostream>
#include <vector>
#include <sstream>
#include <cmath>

typedef std::vector<std::string> vs;
typedef std::vector<int> vi;

int get_minutes(std::string time) {
    return ((time[0]-'0')*10 + (time[1] - '0'))*60 + ((time[3]-'0')*10 + (time[4] - '0'));
}

int minimum_time_difference(vs arr) {
    vi v;
    for (int i=0; i<arr.size(); i++) {
        v.push_back(get_minutes(arr[i]));
    }
    sort(v.begin(), v.end());
    v.push_back(v[0]+get_minutes("24:00"));
    int min_val = INT_MAX;
    for (int i=1; i<v.size(); i++) {
        min_val = std::min(min_val, v[i]-v[i-1]);
    }
    return min_val;
}

int main() {
    std::cout<<minimum_time_difference(vs{"23:59","00:00"})<<"\n";
}

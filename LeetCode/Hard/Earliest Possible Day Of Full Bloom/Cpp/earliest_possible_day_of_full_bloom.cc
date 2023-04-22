#include <iostream>
#include <vector>

typedef std::vector<int> vi;
typedef std::pair<int, int> pi;
typedef std::vector<pi> vpi;

bool sort_by_second(const pi &a, const pi &b) { return a.second > b.second; }

int earliest_possible_day_of_full_bloon(vi &plant, vi &grow) {
    vpi v;
    for (int i=0;i<plant.size();i++){
        v.push_back(std::make_pair(plant[i], grow[i]));
    }
    std::sort(v.begin(), v.end(), sort_by_second);
    int res = 0;
    int curr_plant = 0;
    for (int i=0;i<v.size(); i++) {
        curr_plant += v[i].first;
        res = std::max(res, curr_plant + v[i].second);
    }
    return res;
}

int main() {
    vi plant{1,2,3,2};
    vi grow{2,1,2,1};
    std::cout<<earliest_possible_day_of_full_bloon(plant, grow)<<"\n";
}

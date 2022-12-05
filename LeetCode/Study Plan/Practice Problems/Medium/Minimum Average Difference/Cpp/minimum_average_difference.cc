#include <iostream>
#include <vector>

typedef long long ll;
typedef std::vector<ll> vi;

int minimum_average_difference(vi v) {
    vi left(v.size()), right(v.size());
    left[0] = v[0];
    right[v.size()-1] = v[v.size()-1];
    for (int i=1; i<v.size(); i++) left[i] = left[i-1] + v[i];
    for (int i=v.size()-2; i>=0; i--) right[i] = right[i+1] + v[i];
    int avg = right[0]/v.size();
    int ind = v.size()-1;
    int l=1, r=v.size()-1;
    for (int i=1; i<v.size(); i++) {
        if (avg == std::abs(right[i]/r-left[i-1]/l)) ind = std::min(ind, i-1);
        if (avg > std::abs(right[i]/r-left[i-1]/l)) {
            ind = i-1;
            avg = abs(right[i]/r-left[i-1]/l);
        }
        l++;
        r--;
    }
    return ind;
}

int main() {
    std::cout<<minimum_average_difference(vi{2,5,3,9,5,3})<<"\n";
}

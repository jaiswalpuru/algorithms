#include <iostream>
#include <vector>

#define PB push_back

typedef std::vector<int> vi;
typedef std::vector<bool> vb;

vb arithmetic_subarrays(vi arr, vi l, vi r) {
    vb res(l.size());
    for (int i=0; i<l.size(); i++) {
        vi temp;
        if ((r[i]-l[i]) < 1 ) continue;
        for (int j=l[i]; j<=r[i]; j++) {
            temp.PB(arr[j]);
        }
        res[i] = true;
        sort(temp.begin(), temp.end());
        for (int j=temp.size()-1; j>=2; j--) {
            if (temp[j]-temp[j-1] != temp[j-1]-temp[j-2]){
                res[i] = false;
                break;
            }
        }
    }
    return res;
}

int main() {
    vb res = arithmetic_subarrays(vi{4,6,5,9,3,7}, vi{0,0,2}, vi{2,3,5});
    for (auto it=res.begin(); it!=res.end(); it++) {
        std::cout<<*it<<" ";
    }
    std::cout<<"\n";
}

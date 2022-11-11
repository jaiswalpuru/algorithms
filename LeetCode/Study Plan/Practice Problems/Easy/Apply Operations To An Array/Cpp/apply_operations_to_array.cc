#include <iostream>
#include <vector>

typedef std::vector<int> vi;

template<typename T>
std::ostream& operator<<(std::ostream& os, const std::vector<T>& v) {
    os<<"[";
    for (int i=0;i<v.size();i++) {
        os<<v[i]<<" ";
    }
    os<<"]\n";
    return os;
}

vi apply_operations_to_array(vi v) {
    for (int i=0; i<v.size()-1; i++) {
        if (v[i] == v[i+1]) {
            v[i] *= 2;
            v[i+1] = 0;
        }
    }
    int i=0, j=0;
    while(i<v.size() && j<v.size()) {
        if (v[j] != 0) {
            v[i] = v[j];
            i++;
            j++;
        }else {
            j++;    
        }
    }
    for (int k=i;k<v.size();k++){
        v[k] = 0;
    }
    return v;
}

int main() {
    std::cout<<apply_operations_to_array(vi{1,2,2,1,1,0});
}

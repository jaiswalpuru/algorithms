#include <iostream>

typedef std::string st;

st max_value_after_insertion(st n, int x) {
    int len = n.size()-1;
    int pos = len+1;
    if (n[0]=='-') {
        for(int i=len;i>=1;i--){
            if (n[i]-'0' > x) pos = i;
        }
    }else {
        for(int i=len;i>=0;i--){
            if (n[i]-'0' < x) pos = i;
        }
    }
    return n.substr(0, pos) + std::to_string(x) + n.substr(pos, len+1);
}

int main() {
    std::cout<<max_value_after_insertion("99", 9)<<"\n";
}

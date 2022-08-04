#include <iostream>
#include <map>

typedef unsigned long long ll;

void missing_number(int n) {
    ll val;
    std::map<int, bool> visited;
    for (int i=0; i<n-1; i++){
        std::cin>>val;
        visited[val] = true;
    }

    for (int i=1; i<=n; i++){
        if (!visited[i]){
            std::cout<<i<<"\n";
            break;
        }
    }
}

int main() {
    int n;
    std::cin>>n;
    missing_number(n);
    return 0;
}



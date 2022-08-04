#include <iostream>

using namespace std;

void permutations(int n){
    if (n == 1) {
        cout<<n<<"\n";
        return;
    }

    if (n <= 3) {
        cout<<"NO SOLUTION\n";
        return;
    }
    
    for (int i=n; i>=1; i--){
        if (i%2 == 1){
            cout<<i<<" ";
        }
    }
    for (int i=n; i>=1; i--){
        if (i%2 == 0){
            cout<<i<<" ";
        }
    }
    cout<<"\n";
} 

int main() {
    int n;
    cin>>n;
    permutations(n);
    return 0;
}


#include <iostream>
#include <vector>
#include <map>

using namespace std;

void generate_subsequences(vector<int> arr, int ind, int n, vector<vector<int> > &res, map<vector<int>, int> &mp, 
vector<int> &temp){
    if (ind >= n){
        if (temp.size()>1){
            if (mp.find(temp) == mp.end()) {
                mp[temp]++;
                res.push_back(vector<int>(temp));
            }
        }
        return ;
    } 

    generate_subsequences(arr, ind+1, n, res, mp, temp);

    if ((temp.size() == 0) || (arr[ind] >= temp[temp.size()-1]) ) {
        temp.push_back(arr[ind]);
        generate_subsequences(arr,ind+1, n, res, mp, temp);
        temp.pop_back();
    }
}

vector<vector<int> > find_subsequences(vector<int> &arr){
    if (arr.size()==0){
        return vector<vector<int> >();
    }
    map <vector<int>, int> mp;
    vector<vector<int> > res;
    vector<int> temp;
    int n = arr.size();
    generate_subsequences(arr, 0, n, res, mp, temp);
    return res;
}


int main() {
    vector<int> arr;
    arr.push_back(4);
    arr.push_back(6);
    arr.push_back(7);
    arr.push_back(7);
    vector<vector<int> > res;
    res = find_subsequences(arr);
    vector<vector<int> >::iterator r;
    vector<int>::iterator c;

    for (r = res.begin(); r != res.end(); r++){
        for (c = r->begin(); c!= r->end(); c++){
            cout<<*c<<" ";
        }
        cout<<endl;
    }


    return 0;
}
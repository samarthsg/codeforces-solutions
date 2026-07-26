#include<iostream>
#include<vector>
using namespace std;


bool checkforpossiblity(vector<int>& d, int n){
    if((d[0] & 1) == 1 && (d[n-1] & 1) == 1 && (n&1) == 1) return true;
    return false;
}
int main(){
    int n;
    cin >> n;
    vector<int> d(n);
    for(int i = 0; i < n; i++){
        cin >> d[i];
    }

    bool a = checkforpossiblity(d, n);
    cout << (a ? "Yes" : "No") << endl;
    return 0;
}

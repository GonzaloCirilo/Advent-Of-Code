#include <iostream>
#include <vector>
#include <algorithm>
#include <math.h>
using namespace std;

int main(){
    int n; char c;
    vector<int> arr;
    while(scanf("%d,", &n) != EOF) {
        arr.push_back(n);
    }
    sort(arr.begin(), arr.end());

    long long ans = 1e14;
    for(int i = arr[0]; i < arr[arr.size() - 1]; i++){
        long long aux = 0; 
        for(int j = 0; j < arr.size(); j++) {
            int x = abs(arr[j] - i);
            aux += x*(x+1) * 1/2;
        }
        ans = min(aux,ans);
    }
    cout << ans;
    return 0;
}
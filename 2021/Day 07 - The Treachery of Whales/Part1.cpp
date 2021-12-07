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

    int mid = arr.size() / 2;

    int ans = 0;
    for(int i = 0; i < arr.size(); i++) {
        ans += abs(arr[i] - arr[mid]);
    }
    cout << ans;
    return 0;
}
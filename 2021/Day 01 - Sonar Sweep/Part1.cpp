#include <iostream>
using namespace std;

int main(){
    int n, prev = 1e9, ans = 0;
    while(cin >> n) {
        if(n > prev)
            ans++;
        prev = n;
    }
    cout << ans;
    return 0;
}
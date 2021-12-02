#include <iostream>
using namespace std;

int main(){
    int x2, x1, x0, prevSum = 1e9, ans = 0, n;
    cin >> x0 >> x1 >> x2;
    prevSum = x0 + x1 + x2;
    while(cin >> n) {
        x0 = x1; x1 = x2; x2 = n;
        if((x0 + x1 + x2) > prevSum)
            ans++;
        prevSum = (x0 + x1 + x2);
    }
    cout << ans;
    return 0;
}
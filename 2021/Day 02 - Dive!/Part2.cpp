#include <iostream>
#include <string>
using namespace std;

int main(){
    string cmd; int n;
    int h = 0, d = 0, a = 0; long long ans = 0;
    while(cin >> cmd >> n) {
        if(cmd[0] == 'f'){
            h += n;
            d += a * n;
        }else if(cmd[0] == 'd') {
            a += n;
        }else {
            a -= n;
        }
    }
    cout << 1LL * h * d;
    return 0;
}
#include <iostream>
using namespace std;

int main(){
    int lx, rx, ly, uy;
    scanf("target area: x=%d..%d, y=%d..%d", &lx, &rx, &ly, &uy);
    printf("target area: x=%d..%d, y=%d..%d\n", lx, rx, ly, uy);
    int vy, vx = 3;
    while(true){
        int r = vx * (vx + 1) / 2;
        if(r >= lx && r <= rx){
            break;
        }
        vx++;
    }
    vy = 100;
    bool found = false;
    while(true){
        int r = vy * (vy + 1) / 2;
        int s = -(vy + 1);
        while(true){
            if(s >= ly && s <= uy){
                found = true;
                break;
            }
            if(s < ly){
                break;
            }
            s += (s - 1);
        }
        if(found) {
            cout << r << endl;;
            break;
        }
        vy--;
    }
    cout << vx << " " << vy;;
    return 0;
}
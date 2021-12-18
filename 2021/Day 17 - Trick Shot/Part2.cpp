#include <iostream>
#include <math.h>
using namespace std;

int main(){
    int lx, rx, ly, uy;
    scanf("target area: x=%d..%d, y=%d..%d", &lx, &rx, &ly, &uy);
    printf("target area: x=%d..%d, y=%d..%d\n", lx, rx, ly, uy);
    int ans = ((rx - lx) + 1) * ((abs(ly) - abs(uy)) + 1);
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
    for(int i = vx; i <= rx / 2; i++){
        for(int j = uy + 1; j <= vy; j++){
            auto [sx, sy] = pair<int,int>(0,0);
            auto dx = i, dy = j;
            while(true){
                if(sx >= lx && sx <= rx && sy >= ly && sy <= uy){
                    cout << i << " " << j << endl;
                    ans++;
                    break;
                }

                if(sx > rx || sy < ly) break;
                sx += dx;
                sy += dy;
                dx--;
                if(dx < 0) dx = 0;
                dy--;
            }
        }
    }
    cout << ans;
    return 0;
}
#import <Foundation/Foundation.h> // GNU GCC v7.1.1

int main (int argc, const char * argv[]){
    NSAutoreleasePool * pool = [[NSAutoreleasePool alloc] init];

    int moons[4][3] = {}, m = 0, vel[4][3] = {}, pot[4] = {}, kin[4] = {}, k;

    for(k = 0; k < 4; k++){
        scanf("<x=%d, y=%d, z=%d>\n", &moons[k][0], &moons[k][1], &moons[k][2]);
    }

    m = 0;
    while(m < 1000){
        int i;
        for(i = 0; i < 3; i++){
            int j;
            for(j = 0; j < 4; j++){
                int k;
                for(k = 0; k < 4; k++){
                    if(j == k) continue;
                    if(moons[k][i] > moons[j][i]){
                        vel[j][i]++;
                    }else if(moons[k][i] < moons[j][i]){
                        vel[j][i]--;
                    }
                }

            }
        }
        for(i = 0; i < 3; i++){
            int j;
            for(j = 0; j < 4; j++){
                moons[j][i] += vel[j][i];
            }
        }
        m++;
    }
    int ans = 0;
    int i;
    for(i = 0; i < 4; i++){
        int j;
        for(j = 0; j < 3; j++){
            pot[i] += abs(moons[i][j]);
            kin[i] += abs(vel[i][j]);
        }
        ans += pot[i] * kin[i];
    }

    printf("%d", ans);

    [pool drain];
    return 0;
}
#import <Foundation/Foundation.h> // GNU GCC v7.1.1

int getVelValue(int a, int b){
    if(a > b) {
        return 1;
    }else if(a < b) {
        return -1;
    }
    return 0;
}

unsigned long long gcd(unsigned long long a, unsigned long long b){
    while(true){
        if(!a)
            return b;
        b %= a;
        if(!b)
            return a;
        a %= b;
    }
}

unsigned long long lcm(unsigned long long a, unsigned long long b){
    unsigned long long aux = gcd(a, b);
    return aux ? (a * b / aux) : 0; 
}

int main (int argc, const char * argv[]){
    NSAutoreleasePool * pool = [[NSAutoreleasePool alloc] init];

    int moons[4][3] = {}, vel[4][3] = {}, k;

    for(k = 0; k < 4; k++){
        scanf("<x=%d, y=%d, z=%d>\n", &moons[k][0], &moons[k][1], &moons[k][2]);
    }

    unsigned long long ans[3] = {};
    int j;
    for(j = 0; j < 3; j++){
        unsigned long long m = 0;
        int aux[8] = {};
        int i;
        for(i = 0; i < 4; i++){
            aux[i] = moons[i][j];
        }
        while(true){
            if(aux[0] == moons[0][j] && aux[1] == moons[1][j] && aux[2] == moons[2][j] && aux[3] == moons[3][j] &&
                aux[4] == vel[0][j] && aux[5] == vel[1][j] && aux[6] == vel[2][j] && aux[7] == vel[3][j] && m != 0)
                break;
            int l;
            for(l = 0; l < 4; l++){
                int n;
                for(n = 0; n < 4; n++){
                    if (l == n) continue;
                    aux[l + 4] += getVelValue(aux[n], aux[l]);
                }
            }
            for(l = 0; l < 4; l++){
                aux[l] += aux[l + 4]
            }
            m++;
        }
        ans[j] = m;
    }

    printf("%llu", lcm(lcm(ans[0], ans[1]), ans[2]));

    [pool drain];
    return 0;
}
#include <iostream>
#include <string>
#include <sstream>
#include <map>
#include <vector>
#include <algorithm>
using namespace std;
unsigned long long a = 0, b = 0;
map<int, int>  ap;

void rec(int score1, int score2, int s1, int s2, unsigned long long sols, int t) {
    if(score1 >= 21){
        a += sols;
        return;
    }
    if(score2 >= 21){
        b += sols;
        return;
    }

    for(auto [k1,v1]: ap){
        if(t == 1){
            auto ns1 = ((k1 + s1 - 1) % 10) + 1;
            rec(score1 + ns1, score2, ns1, s2, sols * v1, t * -1);
        }else{
            auto ns2 = ((k1 + s2 - 1) % 10) + 1;
            rec(score1, score2 + ns2, s1, ns2, sols * v1, t * -1);
        }
    }
}

int main(){
    for(int i = 1; i <= 3; i++){
        for(int j = 1; j <= 3; j++){
            for(int k = 1; k <= 3; k++){
                ap[i+j+k]++;
            }
        }
    }

    int s1, s2, d;
    scanf("Player %d starting position: %d\n", &d, &s1);
    scanf("Player %d starting position: %d\n", &d, &s2);
    rec(0, 0, s1, s2, 1LL,1);
    cout << max(a, b) << " " << min(a,b);

    return 0;
}
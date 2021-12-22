#include <iostream>
#include <string>
#include <sstream>
using namespace std;

void solve(){
    int s1, s2, d;
    scanf("Player %d starting position: %d\n", &d, &s1);
    scanf("Player %d starting position: %d\n", &d, &s2);
    int score1 = 0, score2 = 0, count = 1, roll = 0;;

    while(score1 < 1000 && score2 < 1000){
        int aux = 0;
        for(int i = 0; i < 3; i++){
            aux += count++;
            if(count == 101)
                count = 1;
            roll++;
            
        }
        s1 = ((aux + s1 - 1) % 10) + 1;
        score1 += s1;
        if(score1 >= 1000){
            cout << 1LL * score2 * roll;
            return;
        }
        aux = 0;
        for(int i = 0; i < 3; i++){
            aux += count++;
            if(count == 101)
                count = 1;
            roll++;
        }
        s2 = ((aux + s2 - 1) % 10) + 1;
        score2 += s2;
        if(score2 >= 1000){
            cout << 1LL * score1 * roll;
            return;
        }
    }
}

int main(){
    solve();
    return 0;
}
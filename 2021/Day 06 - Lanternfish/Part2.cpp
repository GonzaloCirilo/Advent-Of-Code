#include <string>
#include <iostream>
using namespace std;

int main() {
    string s;
    getline(cin, s);
    long long days[9] = {}, aux[9] = {}, ans = 0;
    for(int i = 0; i < s.size(); i++) {
        if(s[i] == ',') continue;
        days[s[i] - '0'] += 1;
    }

    for(int i = 0; i < 256; i++) {
        for(int j = 0; j < 9; j++) {
            if(j == 0) {
                aux[6] += days[j];
                aux[8] += days[j];
                days[j] = 0;
            }else{
                aux[j - 1] += days[j];
                days[j] = 0;
            }
        }
        for(int j = 0; j < 9; j++) {
            days[j] += aux[j];
            aux[j] = 0;
        }
    }
    for(int j = 0; j < 9; j++) {
        ans += days[j];
    }
    cout << ans;
    return 0;
}
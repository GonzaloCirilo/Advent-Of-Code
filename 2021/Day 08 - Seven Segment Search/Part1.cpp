#include <iostream>
#include <string>
#include <vector>
#include <sstream>
using namespace std;

int main(){
    string s;
    cout << "ans: ";
    int ans = 0;
    while(getline(cin, s)){
        stringstream ss(s);
        string aux; bool passed = false;
        while(ss >> aux){
            if (passed) {
                ans += (int)(aux.size() == 2 || aux.size() == 4 || aux.size() == 3 || aux.size() == 7);
            }
            if(aux == "|"){
                passed = true;
            }
        }
    }
        cout << ans;
    return 0;
}
#include <iostream>
#include <string>
#include <vector>
using namespace std;

int main(){
    string l; 
    vector<string>lines;
    int n[15]  = {}, count = 0;
    while(cin >> l) {
        lines.push_back(l);
        count++;
    }

    long long g = 0, e = 0;
    for(int i = 0; i < l.size(); i++) {
        for(int j = 0; j < lines.size(); j++)
            n[i] += (lines[j][i] == '1');
        //cout << count << " " << n[i] << endl;
        if (n[i] > (count / 2)){
            g += (1 << (l.size() - 1 - i));
        } else {
            e += (1 << (l.size() - 1 - i));
        }
    }

    cout << 1LL * g * e;
    return 0;
}
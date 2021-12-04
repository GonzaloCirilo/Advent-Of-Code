#include <iostream>
#include <string>
#include <vector>
using namespace std;

int main(){
    string l; 
    vector<string>lines, aux, copy;
    int n[15]  = {}, n2[15] = {}, count = 0;
    while(cin >> l) {
        lines.push_back(l);
        count++;
    }

    long long o = 0, c = 0;
    copy = lines;
    for(int i = 0; i < l.size(); i++) {
        if(lines.size() == 1){
            c = 0;
            for(int j = 0; j < l.size(); j++){
                c += ((lines[0][j] - '0') << (l.size() - 1 - j));
            }
            break;
        }
        for(int j = 0; j < lines.size(); j++){
            n[i] += (lines[j][i] == '1');
        }
        if (n[i] >= count - n[i]){
            for(int j = 0; j < lines.size(); j++){
                if(lines[j][i] == '0'){
                    aux.push_back(lines[j]);
                }
            }
        }else {
            c += (1 << (l.size() - 1 - i));
            for(int j = 0; j < lines.size(); j++){
                if(lines[j][i] == '1'){
                    aux.push_back(lines[j]);
                }
            }
        }
        lines = aux;
        count = aux.size();     
        aux.clear();
    }
    lines = copy;
    count = copy.size();
    for(int i = 0; i < l.size(); i++) {
        if(lines.size() == 1){
            o = 0;
            for(int j = 0; j < l.size(); j++){
                o += ((lines[0][j] - '0') << (l.size() - 1 - j));
            }
            break;
        }
        for(int j = 0; j < lines.size(); j++){
            n2[i] += (lines[j][i] == '1');
        }
        if (n2[i] >= count - n2[i]){
            o += (1 << (l.size() - 1 - i));
            for(int j = 0; j < lines.size(); j++){
                if(lines[j][i] == '1'){
                    aux.push_back(lines[j]);
                }
            }
        }else {
            for(int j = 0; j < lines.size(); j++){
                if(lines[j][i] == '0'){
                    aux.push_back(lines[j]);
                }
            }
        }
        lines = aux;
        count = aux.size();  
        aux.clear();   
    }
    cout << o << " " << c << endl;
    cout << 1LL * o * c;
    return 0;
}
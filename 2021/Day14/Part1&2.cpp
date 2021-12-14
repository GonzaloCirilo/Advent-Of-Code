#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <map>
using namespace std;

void printAnswer(long long p[26], int step){
    vector<long long> arr;
    for(int i = 0 ; i < 26; i++){
        if(p[i]!=0){
            arr.push_back(p[i]);
        }
    }
    auto x = *max_element(arr.begin(), arr.end()), y = *min_element(arr.begin(), arr.end());
    cout << "Step " << step << ": " << x - y << endl;
}

int main(){
    string s, pair, f, r;
    cin >> s;
    map<string, string> ma;
    map<string, long long> occ;
    long long auxOC[26] = {};
    for(int i = 0; i < s.size() - 1; i++){
        auxOC[s[i] - 'A']++;
        if(!occ.insert({s.substr(i, 2), 1}).second){
            occ[s.substr(i,2)]++;
        }
    }
    auxOC[s[s.size() - 1] - 'A']++;
    
    while(cin >> pair >> f >> r){
        ma[pair] = r;
    }
    int step = 0;
    while(step != 40){
        if(step == 10){ printAnswer(auxOC, 10); }
        auto auxOccur = occ;
        for(auto it = occ.begin(); it != occ.end(); it++){
            string token = (*it).first; long long occur = (*it).second;
            if ( ma.find(token) != ma.end() ) {
                string inserted = ma[token], newL = token[0] + inserted, newR = inserted + token[1];
                auxOccur[token] -= occur;
                auxOccur[newL] += occur;
                auxOccur[newR] += occur;
                auxOC[inserted[0] - 'A'] += occur;
            }
        }
        occ = auxOccur;
        step++;
    }
    printAnswer(auxOC, 40);
    return 0;
}
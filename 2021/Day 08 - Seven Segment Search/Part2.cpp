#include <iostream>
#include <string>
#include <vector>
#include <set>
#include <sstream>
#include <algorithm>
using namespace std;

int main(){
    string s;
    
    int ans = 0;
    while(getline(cin, s)){
        stringstream ss(s);
        string aux; bool passed = false;
        vector<string> tokens;
        while(ss >> aux){
            tokens.push_back(aux);
        }
        sort(tokens.begin(), tokens.end(), []
            (const std::string& first, const std::string& second){
                return first.size() < second.size();
            });

        set<char> fixed[2] = {};
        for(int i = 0; i < tokens.size(); i++){
            string aux = tokens[i];
            int num = 0;
            if(aux.size() == 2) {
                fixed[0].insert(aux[0]);
                fixed[0].insert(aux[1]);
            }else if (aux.size() == 4) {
                for(int j = 0; j < aux.size(); j++) {
                    auto cd = fixed[0];
                    if(cd.insert(aux[j]).second){
                        fixed[1].insert(aux[j]);
                    }
                }
            }
        }

        stringstream news(s);
        int num = 0;
        while(news >> aux){
            if (passed) {
                int d = -1;
                if(aux.size() == 7){
                    d = 8;
                }else if(aux.size() == 4) {
                    d = 4;
                }else if(aux.size() == 3) {
                    d = 7;
                }else if (aux.size() == 2) {
                    d = 1;
                }else if (aux.size() == 6){
                    auto l = fixed[1], t = fixed[0];
                    for(int i = 0; i < aux.size(); i++) {
                        t.erase(aux[i]);
                        l.erase(aux[i]);
                    }
                    if(t.size() == 0 && l.size() == 0){
                        d = 9;
                    }else if(t.size() == 1){
                        d = 6;
                    }else{
                        d = 0;
                    }
                }else if(aux.size() == 5){
                    auto l = fixed[1], t = fixed[0];
                    for(int i = 0; i < aux.size(); i++) {
                        t.erase(aux[i]);
                        l.erase(aux[i]);
                    }
                    if(t.size() == 1 && l.size() == 1){
                        d = 2;
                    }else if(l.size() == 1){
                        d = 3;
                    }else{
                        d = 5;
                    }
                }
                num *= 10;
                num += d;
            }
            if(aux == "|"){
                passed = true;
            }
        }
        ans += num;
    }
    cout << "ans: " << ans;
    return 0;
}
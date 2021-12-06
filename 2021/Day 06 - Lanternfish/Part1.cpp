#include <string>
#include <iostream>
#include <vector>
using namespace std;

int main() {
    vector<int> fish = vector<int>(), aux = fish;
    string s;
    getline(cin, s);
    for(int i = 0; i < s.size(); i++) {
        if(s[i] == ',') continue;
        fish.push_back(s[i] - '0');
    }

    for(int i = 0; i < 80; i++) {
        for(int j = 0; j < fish.size(); j++) {
            if(fish[j] == 0) {
                fish[j] = 6;
                aux.push_back(8);
            }else{
                fish[j] -= 1;
            }
        }

        for(int j = 0; j < aux.size(); j++) {
            fish.push_back(aux[j]);
        }
        aux.clear();
    }
    cout << fish.size();
    return 0;
}
#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <map>
#include <stack>
using namespace std;

int main(){
    string line;
    vector<long long> arr;
    map<char, pair<char, int>> m; map<char, int> ms;
    m.insert({')', {'(', 3}});
    m.insert({'}', {'{', 1197}});
    m.insert({']', {'[', 57}});
    m.insert({'>', {'<', 25137}});
    ms.insert({'(', 1});
    ms.insert({'[', 2});
    ms.insert({'{', 3});
    ms.insert({'<', 4});

    while(getline(cin, line)){
        stack<char> s;
        long long score = 0;
        for(int i = 0; i < line.size(); i++) {
            if(line[i] == '{' || line[i] == '(' || line[i] == '<' || line[i] == '[') {
                s.push(line[i]);
            }else{
                auto [expects, localScore] = m[line[i]];
                if(expects != s.top()) {
                    score += localScore;
                    break;
                }
                s.pop();
            }
        }
        if(score == 0) {
            long long c = 0;
            while(!s.empty()) {
                c *= 5;
                char cc = s.top(); s.pop();
                c += ms[cc];
            }
            arr.push_back(c);
        }
    }
    sort(arr.begin(), arr.end());
    cout << arr[arr.size()/2];
    return 0;
}
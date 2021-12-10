#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <map>
#include <stack>
using namespace std;

int main(){
    string line;
    long long ans;
    map<char, pair<char, int>> m;
    m.insert({')', {'(', 3}});
    m.insert({'}', {'{', 1197}});
    m.insert({']', {'[', 57}});
    m.insert({'>', {'<', 25137}});

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
        ans += score;
    }
    cout << ans;//395553
    return 0;
}
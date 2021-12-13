#include <iostream>
#include <vector>
#include <algorithm>
#include <map>
#include <stack>
using namespace std;

int ans = 0;
int nodeCount, sn, en;
map<int, string> dic;
vector<vector<int>> grid = vector<vector<int>>(100, vector<int>());
void rec(int ind, vector<bool> visited, vector<bool> isBig, stack<int> s){
    if(ind == en){
        ans++;
        return;
    }
    for(int i = 0; i < grid[ind].size(); i++){
        int newNode = grid[ind][i];
        if(!visited[newNode]){
            if(!isBig[newNode])
                visited[newNode] = true;
            s.push(newNode);
            rec(newNode, visited, isBig, s);
            if(!isBig[newNode])
                visited[newNode] = false;
            s.pop();
        }
    }
}

int main(){
    string s;
    vector<bool> visited, isBig;
    map<string, int> kv;
    while(getline(cin, s)){
        string aux = "", u, v;
        for(int i = 0; i < s.size(); i++){
            if(s[i] == '-'){
                u = aux;
                if(kv.insert({aux, nodeCount}).second){
                    nodeCount++;
                    dic[nodeCount-1] = aux;
                    visited.push_back(false);
                    if(u != "start" && u != "end" && (u[0] - 'A') >= 0 && (u[0] - 'A') < 26){
                        isBig.push_back(true);
                    }else{
                        isBig.push_back(false);
                    }
                }
                aux = "";
                continue;
            }
            aux.push_back(s[i]);
        }
        if(aux != ""){
            v = aux;
            if(kv.insert({aux, nodeCount}).second){
                nodeCount++;
                dic[nodeCount-1] = aux;
                if(v != "start" && v != "end" && (v[0] - 'A') >= 0 && (v[0] - 'A') < 26){
                isBig.push_back(true);
                }else{
                    isBig.push_back(false);
                }
                visited.push_back(false);
            }
            
        }
        grid[kv[u]].push_back(kv[v]);
        grid[kv[v]].push_back(kv[u]);
    }
    sn = kv["start"], en = kv["end"];
    visited[sn] = true;
    stack<int> ss;
    ss.push(sn);
    rec(sn, visited, isBig,ss);
    cout << ans;
    return 0;
}
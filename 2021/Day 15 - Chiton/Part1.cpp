#include <iostream>
#include <string>
#include <vector>
#include <queue>
#include <sstream>
#include <functional>
using namespace std;
typedef pair<int, int> ii;
typedef pair<int, ii> iii;

bool valid(int x, int y, int maxX, int maxY){
    return x >= 0 && x < maxX && y >= 0 && y < maxY;
}

int main(){
    string s;
    vector<vector<int>> grid, dist;
    int dx[4] = {0,0,1,-1}, dy[4] = {1, -1, 0, 0};
    while(getline(cin, s)){
        stringstream ss(s);
        vector<int> aux, aux3;
        for(int i = 0; i < s.size(); i++){
            aux.push_back(s[i] - '0');
            aux3.push_back(1e9);
        }
        grid.push_back(aux);
        dist.push_back(aux3);
    }

    priority_queue<iii, vector<iii>, greater<iii>> pq;
    pq.push({0, {0,0}});
    while(!pq.empty()){
        auto [cost, pos] = pq.top(); pq.pop();
        for(int i = 0; i < 4; i++){
            int newY = pos.first + dy[i], newX = pos.second + dx[i];
            if(valid(newX, newY, grid.size(), grid[0].size())){
                if(cost + grid[newY][newX] < dist[newY][newX]){
                    pq.push({cost + grid[newY][newX], {newY, newX}});
                    dist[newY][newX] = cost + grid[newY][newX];
                }
            }
        }
    }
    cout << dist[grid.size() - 1][grid[0].size() - 1];
    return 0;
}
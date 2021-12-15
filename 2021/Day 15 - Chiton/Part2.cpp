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
    vector<vector<int>> grid, dist, bigGrid;
    int dx[4] = {0,0,1, -1}, dy[4] = {1,-1,0,0};
    while(getline(cin, s)){
        stringstream ss(s);
        vector<int> aux;
        for(int i = 0; i < s.size(); i++){
            aux.push_back(s[i] - '0');
        }
        grid.push_back(aux);
    }
    vector<int> increment;
    int l = grid.size();
    for(int i = 0; i < l; i++){
        increment.push_back(i);
    }
    bigGrid = vector<vector<int>>(l *5, vector<int>(l * 5, 0));
    dist = vector<vector<int>>(l *5, vector<int>(l * 5, 1e9));
    for(int i = 0; i < 25; i++){
        int offsetX = (i % 5) * l, offsetY = (i / 5) * l;
        for(int j = 0; j < l; j++){
            for(int k = 0; k < l; k++){
                int x = offsetX + k, y = offsetY + j;
                bigGrid[y][x] = (grid[j][k] + increment[i % 5] + (i/5) - 1 + 9) % 9 + 1;
            }
        }
    }

    priority_queue<iii, vector<iii>, greater<iii>> pq;
    pq.push({0, {0,0}});
    while(!pq.empty()){
        auto [cost, pos] = pq.top(); pq.pop();
        if(cost > dist[pos.first][pos.second])continue;
        for(int i = 0; i < 4; i++){
            int newY = pos.first + dy[i], newX = pos.second + dx[i];
            if(valid(newX, newY, bigGrid.size(), bigGrid.size())){
                if(cost + bigGrid[newY][newX] < dist[newY][newX]){
                    pq.push({cost + bigGrid[newY][newX], {newY, newX}});
                    dist[newY][newX] = cost + bigGrid[newY][newX];
                }
            }
        }
    }
    cout << dist[bigGrid.size() - 1][bigGrid[0].size() - 1];
    return 0;
}
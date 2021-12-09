#include <iostream>
#include <sstream>
#include <algorithm>
#include <vector>
#include <set>
#include <queue>
using namespace std;

bool valid(int i, int j, int maxi, int maxj){
    return i >= 0 && i < maxi && j >= 0 && j < maxj;
}

int main() {
    string s;
    vector<vector<int>> grid = vector<vector<int>>();
    int dx[4] = {0,0,1,-1}, dy[4] = {1,-1,0,0};
    while(getline(cin, s)){
        stringstream ss(s);
        char n;
        vector<int> arr;
        while(ss >> n){
            arr.push_back(n-'0');
        }
        grid.push_back(arr);
    }

    set<pair<int,int>> lowPoints;
    for(int i = 0; i < grid.size(); i++) {
        for(int j = 0; j < grid[i].size(); j++) {
            int count = 0, maxCount = 0;
            for(int k = 0; k < 4; k++){
                int ni = i + dx[k], nj = j + dy[k];
                if(valid(ni,nj, grid.size(), grid[i].size())){
                    maxCount++;
                    if(grid[i][j] < grid[ni][nj]){
                        count++;
                    }
                }
            }
            if(count == maxCount){
                lowPoints.insert({i,j});
            }
        }
    }
    vector<long long> basins;
    for(auto it = lowPoints.begin(); it != lowPoints.end(); it++){
        auto [si, sj] = *it;
        queue<pair<int,int>> q;
        vector<vector<bool>> visited = vector<vector<bool>>(grid.size(), vector<bool>(grid[0].size(), false));
        visited[si][sj] = true;
        long long count = 1;
        q.push(*it);
        while(!q.empty()){
            auto [i, j] = q.front(); q.pop();
            for(int k = 0; k < 4; k++){
                int ni = i + dx[k], nj = j + dy[k];
                if(valid(ni,nj, grid.size(), grid[i].size()) && !visited[ni][nj] && grid[i][j] < grid[ni][nj] && grid[ni][nj] != 9){
                    q.push({ni, nj});
                    visited[ni][nj] = true;
                    count++;
                }
            }
        }
        basins.push_back(count);
    }

    sort(basins.rbegin(), basins.rend());
    cout << 1LL * basins[0] * basins[1] * basins[2];
    return 0;
}
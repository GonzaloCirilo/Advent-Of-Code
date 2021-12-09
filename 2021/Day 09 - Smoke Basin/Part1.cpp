#include <iostream>
#include <sstream>
#include <algorithm>
#include <vector>
using namespace std;

bool valid(int i, int j, int maxi, int maxj){
    return i >= 0 && i < maxi && j>=0 && j < maxj;
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

    int ans = 0;
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
                ans += grid[i][j] + 1;
            }
        }
    }
    cout << ans;
    return 0;
}
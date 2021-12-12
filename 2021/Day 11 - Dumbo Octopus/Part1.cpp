#include <string>
#include <vector>
#include <iostream>
#include <stack>
using namespace std;

bool isValid(int i, int j, int maxi, int maxj){
    return i >= 0 && i < maxi && j >= 0 && j < maxj;
}

int main() {
    string s;
    vector<vector<int>> grid;
    int dx[8] = {1, 1, 1,  0, 0, -1, -1, -1 }, dy[8] = {1, 0, -1, 1, -1, 1, 0, -1};
    while(getline(cin, s)){
        vector<int> arr;
        for(int i = 0; i < s.size(); i++) {
            arr.push_back(s[i] - '0');
        }
        grid.push_back(arr);
    }
    long long ans = 0;
    for(int i = 0; i < 100; i++) {
        for(int j = 0; j < grid.size(); j++) {
            for(int k = 0; k < grid[j].size(); k++){
                stack<pair<int,int>> ss;
                ss.push({j, k});
                while(!ss.empty()){
                    auto [x, y] = ss.top(); ss.pop();
                    grid[x][y]++;
                    if(grid[x][y] == 10){
                        for(int l = 0; l < 8; l++){
                            int ni = x + dx[l], nj = y + dy[l];
                            if(isValid(ni, nj, grid.size(), grid[0].size()))
                                ss.push({ni, nj});
                        }
                    }
                }
            }
        }

        for(int j = 0; j < grid.size(); j++) {
            for(int k = 0; k < grid[j].size(); k++){
                if(grid[j][k] > 9){
                    grid[j][k] = 0;
                    ans++;
                }
            }
        }
    }
    cout << ans;
    return 0;
}
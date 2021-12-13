#include <string>
#include <iostream>
#include <vector>
#include <algorithm>
#include <sstream>
using namespace std;

int main(){
    int x, y, maxX = -1, maxY = -1;
    vector<int> xs, ys;
    string s;
    while(getline(cin, s) && s != ""){
        stringstream ss(s);
        char c;
        ss >> x >> c >> y;
        maxX = max(x, maxX);
        maxY = max(y, maxY);
        xs.push_back(x);
        ys.push_back(y);
    }
    vector<vector<int>> grid = vector<vector<int>>(maxY + 3, vector<int>(maxX + 3, 0));
    for(int i = 0; i < xs.size(); i++){
        grid[ys[i]][xs[i]] = 1;
    }
    char c, d; int f, count = 0, ans = 0;
    while(cin >>s >> s >> c >> d >> f) {
        if(c == 'x') {
            for(int i = 0; i <= maxY; i++){
                for(int j = 1; j <= f; j++){
                    grid[i][f-j] += grid[i][f+j];
                }
            }
            maxX = f - 1;
        }else{
            for(int i = 0; i <= maxX; i++){
                for(int j = 1; j <= f; j++){
                    grid[f - j][i] += grid[f + j][i];
                }
            }
            maxY = f-1;
        }
        if (!count) {
            for(int i = 0; i <= maxY; i++){
                for(int j = 0; j <= maxX; j++){
                    ans += (int)(grid[i][j] > 0);
                }
            }
            count++;
        }
    }
    for(int i = 0; i <= maxY; i++){
        for(int j = 0; j <= maxX; j++){
            if(grid[i][j] > 0) {
                cout << "#";
            }else{
                cout << ".";
            }
        }
        cout <<endl;
    }
    cout << "Part 1: " << ans;
    return 0;
}
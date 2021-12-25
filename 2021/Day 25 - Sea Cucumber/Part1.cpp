#include <iostream>
#include <string>
#include <vector>
using namespace std;

int main(){
    string s;
    vector<string> grid;
    while(cin >> s){
        grid.push_back(s);
    }
    int count = 0;
    while(true){
        int moves = 0;
        auto copy = grid;
        //east most
        for(int i = grid.size() - 1; i >= 0; i--){
            for(int j = grid[0].size() - 1; j >= 0; j--){
                if(grid[i][j] == '>'){
                    int newJ = (j + 1) % grid[0].size();
                    if(grid[i][newJ] == '.'){
                        moves++;
                        copy[i][newJ] = '>';
                        copy[i][j] = '.';
                    }
                }
            }
        }
        grid = copy;
        //south most
        for(int i = grid.size() - 1; i >= 0; i--){
            for(int j = grid[0].size() - 1; j >= 0; j--){
                if(grid[i][j] == 'v'){
                    int newi = (i + 1) % grid.size();
                    if(grid[newi][j] == '.'){
                        moves++;
                        copy[newi][j] = 'v';
                        copy[i][j] = '.';
                    }
                }
            }
        }

        grid = copy;
        if(moves == 0)
            break;
        count++;
    }
    for(int i = 0; i < grid.size(); i++){
        cout << grid[i] << endl;
    }
    cout << count + 1;
    return 0;
}
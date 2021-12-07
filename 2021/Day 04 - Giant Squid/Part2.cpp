#include <iostream>
#include <string>
#include <algorithm>
#include <sstream>
#include <vector>
using namespace std;

bool checkWinnerBoard(vector<bool> marks) {
    int cols[5] = {}, rows[5] = {};
    for(int i = 0; i < 5; i ++) {
        for(int j = 0; j < 5; j++) {
            int ind = i * 5 + j;
            rows[i] += marks[ind];
        }
        if(rows[i] == 5){
            return true;
        }
        for(int j = 0; j < 5; j++) {
            int ind = i + j * 5;
            cols[i] += marks[ind];
        }
        if(cols[i] == 5){
            return true;
        }
    }
    return false;
}

int main(){
    string l; int n; char c;
    vector<int> arr;
    while(scanf("%d%c", &n, &c) && c != '\n') {
        arr.push_back(n);
    }

    vector<vector<int>> boards = vector<vector<int>>();
    vector<vector<bool>> marks = vector<vector<bool>>();
    int x, b = 0, cont = 0;
    vector<int> board = vector<int>();
    while(cin >> x){
        board.push_back(x);
        cont++;
        if (cont == 25) {
            cont = 0;
            boards.push_back(board);
            marks.push_back(vector<bool>(25, false));
            board.clear();
        }
    }

    vector<int> won = vector<int>(boards.size(), 0);

    for(int i = 0; i < arr.size(); i++) {
        int n = arr[i];

        for(int j = 0; j < boards.size(); j++){
            vector<int>::iterator p = find(boards[j].begin(), boards[j].end(), n);
            if (p != boards[j].end()){
                int ind = p - boards[j].begin();
                marks[j][ind] = true;
            }
            if(checkWinnerBoard(marks[j]) && !won[j]){
                int sum = 0;
                for(int k = 0; k < marks[j].size(); k++){
                    if(!marks[j][k]){
                        sum += boards[j][k];
                    }
                }
                won[j] = sum * n;
                cout << sum * n << endl;
            }
        }
    }
    return 0;
}
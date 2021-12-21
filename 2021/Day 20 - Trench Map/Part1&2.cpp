#include <iostream>
#include <string>
#include <sstream>
#include <vector>
using namespace std;

string s, lookup;

char expandedChar(int pass){
    return lookup[0] == '#' ? (pass % 2 == 0? '.' : '#') : '.';
}

char checkLit(int i, int j, vector<string> image, int pass){
    if(i >= 0 && i < image.size() && j >= 0 && j < image[0].size())
        return image[i][j];
    else
        return expandedChar(pass);
}

vector<string> enhance(vector<string> image, int pass){
    auto canvas = vector<string>(image.size(), string(image[0].size(), '.'));
    for(int i = 0; i < image.size(); i++){
        for(int j = 0; j < image[i].size(); j++){
            int num = 0;
            for(int r = -1; r < 2; r++){
                for(int c = -1; c < 2; c++){
                    num <<= 1;
                    num |= (checkLit(i + r, j + c, image, pass) == '#'? 1 : 0);
                }
            }
            canvas[i][j] = lookup[num];
        }
    }
    return canvas;
}

int main(){
    cin >> lookup;
    vector<string> image;
    while(cin >> s){
        image.push_back(s);
    }
    for(int i = 0; i < 50; i++){
        if(i == 2){
            int ans = 0;
            for(int i = 0; i < image.size(); i++){
                for(int j = 0; j < image[i].size(); j++){
                    ans += (int)(image[i][j] == '#');
                }
            }
            cout << ans<< endl;
        }
        for(int j = 0; j < image.size(); j++){
            image[j].insert(0, string(1,expandedChar(i)));
            image[j].push_back(expandedChar(i));
        }
        image.insert(image.begin(), 1, string(image[0].size(), expandedChar(i)));
        image.push_back(string(image[0].size(), expandedChar(i)));
        image = enhance(image, i);

    }
    cout << "Result " << endl;
    int ans = 0;
    for(int i = 0; i < image.size(); i++){
        for(int j = 0; j < image[i].size(); j++){
            ans += (int)(image[i][j] == '#');
        }
    }
    cout << ans;
    return 0;
}
#include <string>
#include <iostream>
#include <list>
#include <algorithm>
#include <vector>
using namespace std;
int sumVersion = 0;
long long toDecimal(vector<int> arr, int &currentInd, int l){
    long long num = 0;
    for(int i = 0; i < l; i++){
        num <<= 1;
        num |= arr[currentInd];
        currentInd++;
    }
    return num;
}

pair<int,unsigned long long> parse(vector<int> arr, int ind){
    int it = ind;;
    sumVersion += toDecimal(arr, it, 3);
    int type = toDecimal(arr, it, 3);
    if(type == 4){
        bool ended = 1;
        long long num = 0;
        do{
            num <<= 4;
            int prefix = arr[it];
            it++;
            ended = prefix;
            num |= toDecimal(arr, it, 4);
        }while(ended);
        return {it, num};
    }else{
        int label = arr[it];
        it++;
        vector<long long> literals;
        if(label){
            int l = toDecimal(arr, it, 11);
            for(int k = 0; k < l; k++) {
                auto [ind, x] = parse(arr, it);
                it = ind;
                literals.push_back(x);
            }
        }else{
            int l = toDecimal(arr, it, 15);
            auto target = it + l;
            while(target != it){
                auto [ind, x] = parse(arr, it);
                it = ind;
                literals.push_back(x);
            }
        }
        unsigned long long res;
        switch (type)
        {
        case 0:
            for(int i = 0; i < literals.size(); i++)
                res += literals[i];
            break;
        case 1:
            res = 1;
            for(int i = 0; i < literals.size(); i++)
                res *= literals[i];
            break;
        case 2:
            res = *min_element(literals.begin(), literals.end());
            break;
        case 3:
            res = *max_element(literals.begin(), literals.end());
            break;
        case 5:
            res = literals[0] > literals[1] ? 1 : 0;
            break;
        case 6:
            res = literals[0] < literals[1] ? 1 : 0;
            break;
        case 7:
            res = literals[0] == literals[1] ? 1 : 0;
            break;
        default:
            break;
        }
        return {it, res};
    }
}

int main(){
    string s;
    cin >> s;
    vector<int> arr;
    for(int i = 0; i < s.size(); i++){
        int d = s[i] - '0';
        if(s[i] >= 'A'){
            d = s[i] - 'A' + 10;
        }
        list<int> aux;
        while (d) {
            aux.push_front(d & 1);
            d >>= 1;
        }
        while(aux.size() != 4) {
            aux.push_front(0);
        }
        for(auto j = aux.begin(); j != aux.end(); j++){
            arr.push_back(*j);
        }
    }
    long long ans = parse(arr, 0).second;
    cout << "Part 1: " << sumVersion << endl << "Part 2: " << ans;
    return 0;
}
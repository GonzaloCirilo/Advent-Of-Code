#include <iostream>
#include <math.h>
using namespace std;

int fuelReq(int n){
	int ans = 0;
	while(n != 0){
		n = (int)floor((double) n / 3) - 2;
		n = n > 0 ? n : 0;
		ans += n;
	}
	return ans;
}

int main() {
	// your code goes here
	int n, fuel;
	while(cin >> n){
		//fuel += (int)floor((double) n / 3) - 2;
		fuel += fuelReq(n);
	}
	cout << fuel;
	return 0;
}
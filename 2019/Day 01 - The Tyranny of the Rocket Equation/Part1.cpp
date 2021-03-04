#include <iostream>
#include <math.h>
using namespace std;

int main() {
	// your code goes here
	int n, fuel;
	while(cin >> n){
		fuel += (int)floor((double) n / 3) - 2;
	}
	cout << fuel;
	return 0;
}
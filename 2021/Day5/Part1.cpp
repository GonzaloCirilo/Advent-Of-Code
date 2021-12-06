#include <iostream>
#include <set>
#include <algorithm>

using namespace std;

int main() {
    int x1, y1, x2, y2;
    set<pair<int, int>> s, saux;
    while (scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2)!= EOF) {
        if (x1 == x2) {
            int minE = min(y1, y2), maxE = max(y1, y2);
            for (int i = minE; i <= maxE; i++) {
                if (!s.insert({ x1, i }).second) {
                    saux.insert({ x1, i });
                }
            }
        }
        else if (y1 == y2) {
            int minE = min(x1, x2), maxE = max(x1, x2);
            for (int i = minE; i <= maxE; i++) {
                if (!s.insert({ i, y1 }).second) {
                    saux.insert({ i, y1 });
                }
            }
        }
    }
    cout << saux.size();
    return 0;
}
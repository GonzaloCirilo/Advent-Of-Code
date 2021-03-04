#include <stdio.h>

int main(){
    int lb, ub, ans = 0;
    char c;

    while(scanf("%d-%d %c: ", &lb, &ub, &c) != EOF){
        char aux = ' '; int count = 0;
        while(aux != '\n' && aux != EOF){
            count += (aux == c);
            aux = getchar();
        }
        ans += (count >= lb && count <= ub);
    }
    printf("%d\n", ans);
    return 0;
}
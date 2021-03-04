#include <stdio.h>

int main(){
    int lb, ub, ans = 0;
    char c;

    while(scanf("%d-%d %c: ", &lb, &ub, &c) != EOF){
        char aux = ' ', pass[30]; int count = 0;
        while(aux != '\n' && aux != EOF){
            aux = getchar();
            pass[count++] = aux;
        }
        ans += ((pass[lb-1] == c) ^ (pass[ub-1] == c));
    }
    printf("%d\n", ans);
    return 0;
}
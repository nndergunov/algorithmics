#include <stdio.h>
#include <stdlib.h>

int cmp(const void* a, const void* b) {
    if(*(int*)a == *(int*)b) {
        return 0;
    } else if (*(int*)a < *(int*)b) {
        return -1;
    } else {
        return 1;
    }
}

int main() {
    int n;
    int k;

    scanf("%d %d", &n, &k);

    int canisters[n];

    for (int i = 0; i<n; i++) {
        scanf("%d", &canisters[i]);

        if (canisters[i] > k) {
            printf("%s", "Impossible");

            return 0;
        }
    }

    qsort(canisters, n, sizeof(int), cmp);

    int res = 0;
    int i = 0;
    int j = 0;

    while (i+j < n) {
        res++;

        if (i!=n-j-1 && canisters[i]+canisters[n-j-1] <= k) {
            i++;
        }

        j++;
    }

    printf("%d", res);
}
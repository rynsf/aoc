#include <stdio.h>
#include <stdlib.h>

#define MAX_LIST_SIZE 1000

int compar(const void *e1, const void *e2) {
    int a = *((int*)e1);
    int b = *((int*)e2);
    if(a > b) {
        return 1;
    }
    if(a < b) {
        return -1;
    }
    return 0;
}

int main() {
    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    int list1[MAX_LIST_SIZE];
    int list2[MAX_LIST_SIZE];
    int a,b;
    int count = 0;
    int distance = 0;

    fp = fopen("./input", "r");
    if(fp == NULL) {
        return 1;
    }

    while(getline(&line, &len, fp) != -1) {
        if(count >= MAX_LIST_SIZE) {
            return 1;
        }
        sscanf(line, "%d   %d", &a, &b);
        list1[count] = a;
        list2[count] = b;
        count++;
    }

    qsort(list1, count, sizeof(int), compar);
    qsort(list2, count, sizeof(int), compar);

    for(int n=0; n < count; n++) {
        distance += abs(list1[n] - list2[n]);
    }

    printf("%d\n", distance);

    fclose(fp);
    return 0;
}

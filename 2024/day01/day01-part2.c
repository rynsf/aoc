#include <stdio.h>
#include <stdlib.h>

#define MAX_LIST_SIZE 1000

int main() {
    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    int list1[MAX_LIST_SIZE];
    int list2[MAX_LIST_SIZE];
    int a,b;
    int nelem = 0;
    int similarityScore = 0;
    int freq;

    fp = fopen("./input", "r");
    if(fp == NULL) {
        return 1;
    }

    while(getline(&line, &len, fp) != -1) {
        if(nelem >= MAX_LIST_SIZE) {
            return 1;
        }
        sscanf(line, "%d   %d", &a, &b);
        list1[nelem] = a;
        list2[nelem] = b;
        nelem++;
    }

    for(int n=0; n < nelem; n++) {
        freq = 0;
        for(int m=0; m < nelem; m++) {
            if(list1[n] == list2[m]) {
                freq += 1;
            }
        }
        similarityScore += list1[n] * freq;
    }

    printf("%d\n", similarityScore);
    return 0;
}

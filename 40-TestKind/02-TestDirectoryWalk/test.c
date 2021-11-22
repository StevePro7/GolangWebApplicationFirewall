#include <stdio.h>
#include <stdlib.h>

static char **makeCharArray(int size) {
    printf( "makeCharArray\n");
    return calloc(sizeof(char*), size);
}

static void freeCharArray(char **a, int size) {
    int i;
    printf( "freeCharArray\n");
    for (i = 0; i < size; i++)
        free(a[i]);
    free(a);
}

static void setArrayString(char **a, char *s, int n) {
    printf("setArrayString [%d] = '%s' start\n", n, s);
    a[n] = s;
    printf("setArrayString [%d] = '%s' -end-\n", n, s);
}

static void processArrayString(char **array, int size) {
    int i = 0;
    printf("processArrayString start\n");

    for( i = 0; i < size; i++ )
    {
        printf( "code file '%s'\n", ( array[ i ] ) );
    }

    printf("processArrayString -end-\n");
}
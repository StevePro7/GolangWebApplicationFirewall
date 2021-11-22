#include <stdio.h>
#include <stdlib.h>

static char **makeCharArray(int size) {
    fprintf(stderr, "  C makeCharArray %i\n", size);
    return calloc(sizeof(char*), size);
}

static void freeCharArray(char **a, int size) {
    int i;
    fprintf(stderr, "  C freeCharArray %d\n", size);
    for (i = 0; i < size; i++)
        free(a[i]);
    free(a);
}

static void setArrayString(char **a, char *s, int n) {
    fprintf(stderr, "  C setArrayString [%d] = '%s' start\n", n, s);
    a[n] = s;
    fprintf(stderr, "  C setArrayString [%d] = '%s' -end-\n", n, s);
}

static void processArrayString(char **array, int size) {
    int i = 0;
    fprintf(stderr, "  C processArrayString start [%d]\n", i);

    for( i = 0; i < size; i++ )
    {
        fprintf(stderr, "  C code file '%s' [%d]\n", ( array[ i ] ), i );
    }

    fprintf(stderr, "  C processArrayString -end- %i\n", i);
}
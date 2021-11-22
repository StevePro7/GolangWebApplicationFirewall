#include <stdio.h>
#include <stdlib.h>

static char **makeCharArray(int size) {
    fprintf(stderr, "  C makeCharArray\n");
    return calloc(sizeof(char*), size);
}

static void freeCharArray(char **a, int size) {
    int i;
    fprintf(stderr, "  C freeCharArray\n");
    for (i = 0; i < size; i++)
        free(a[i]);
    free(a);
}

static void setArrayString(char **a, char *s, int n) {
    fprintf(stderr, "setArrayString [%d] = '%s' start\n", n, s);
    a[n] = s;
    fprintf(stderr, "setArrayString [%d] = '%s' -end-\n", n, s);
}

static void processArrayString(char **array, int size) {
    int i = 0;
    char *fileName;
    fprintf(stderr, "  C processArrayString start [%d]\n", i);

    for( i = 0; i < size; i++ )
    {
        fileName = array[ i ] ;
        //fprintf(stderr, "  C code file '%s' [%d]\n", ( array[ i ] ), i );
        fprintf(stderr, "  C code file!! '%s'\n", fileName );
    }

    fprintf(stderr, "  C processArrayString -end-\n");
}
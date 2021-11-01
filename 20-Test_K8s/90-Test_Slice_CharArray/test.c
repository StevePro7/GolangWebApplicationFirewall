#include <stdio.h>
#include <stdlib.h>

static char **makeCharArray(int size) {
    return calloc(sizeof(char*), size);
}

static void freeCharArray(char **a, int size) {
    int i;
    for (i = 0; i < size; i++)
        free(a[i]);
    free(a);
}

static void setArrayString(char **a, char *s, int n) {
    printf("ss [%d] = '%s'\n", n, s);
    a[n] = s;
    printf("ee [%d] = '%s'\n", n, s);
}

static void processArrayString(char **array, int size) {
    int i = 0;
    printf("start processArrayString\n");

    for( i = 0; i < size; i++ )
    {
        printf( "%s\n", ( array[ i ] ) );
    }
//    for( i = 0; i < sizeof( array ) / sizeof( array[ 0 ] ); i++ )
//    {
//        printf( "%s\n", ( array[ i ] ) );
//    }

    printf("-end- processArrayString\n");
}
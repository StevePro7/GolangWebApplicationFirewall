#include <stdio.h>
#include <stdlib.h>
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

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

// Function was previously MyCInit()
static void processArrayString(char **array, int size) {
    int i = 0;
    const char *file;
    const char *error = NULL;

    fprintf(stderr, "  C processArrayString start\n");
    if (modsec == NULL)
    {
        fprintf(stderr, "  C msc_init\n");
        modsec = msc_init();

        fprintf(stderr, "  C msc_create_rules_set\n");
        rules = msc_create_rules_set();

        for( i = 0; i < size; i++ )
        {
            file = array[ i ] ;
            fprintf(stderr, "  C code file[%d] start '%s'\n", i, file );
            msc_rules_add_file(rules, file, &error);
            fprintf(stderr, "  C code file[%d] -end- '%s'\n", i, file );
        }
    }

    fprintf(stderr, "  C processArrayString -end-\n");
}
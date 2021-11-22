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
        fprintf(stderr, "  C msc_init()\n");
        modsec = msc_init();

        fprintf(stderr, "  C msc_create_rules_set()\n");
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


int MyCProcess(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port)
{
    Transaction *transaction = NULL;
    transaction = msc_new_transaction(modsec, rules, NULL);

    fprintf(stderr, "  C URI='%s'\n", uri);
    fprintf(stderr, "  C Method='%s'\n", http_method);
    fprintf(stderr, "  C Protocol='%s/%s'\n", http_protocol, http_version);
    fprintf(stderr, "  C Client socket='%s:%d'\n", client_link, client_port);
    fprintf(stderr, "  C Server socket='%s:%d'\n", server_link, server_port);

    fprintf(stderr, "  C msc_process_connection()\n");
    msc_process_connection(transaction, client_link, client_port, server_link, server_port);

    fprintf(stderr, "  C msc_process_uri()\n");
    msc_process_uri(transaction, uri, http_protocol, http_version);

    fprintf(stderr, "  C msc_process_request_headers()\n");
    msc_process_request_headers(transaction);

    fprintf(stderr, "  C msc_process_request_body()\n");
    msc_process_request_body(transaction);

    ModSecurityIntervention intervention;
    intervention.status = 200;
    intervention.url = NULL;
    intervention.log = NULL;
    intervention.disruptive = 0;
    int inter = msc_intervention(transaction, &intervention);

    fprintf(stderr, "intervention=%i\n", inter);
    fprintf(stderr, "inter status=%i\n", intervention.status);
    return inter;
}
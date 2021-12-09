#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

void InitModSec()
{
    const char *error = NULL;
    modsec = msc_init();
    rules = msc_create_rules_set();
    msc_rules_add_file(rules, "/etc/waf/crs-setup.conf", &error);
    msc_rules_add_file(rules, "/etc/waf/modsecdefault.conf", &error);
    msc_rules_add_file(rules, "/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
}

int ProcessHttpRequest(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port)
{
    Transaction *transaction = NULL;
    transaction = msc_new_transaction(modsec, rules, NULL);
    msc_process_connection(transaction, client_link, client_port, server_link, server_port);
    msc_process_uri(transaction, uri, http_protocol, http_version);
    msc_process_request_headers(transaction);
    msc_process_request_body(transaction);

    ModSecurityIntervention intervention;
    intervention.status = 200;
    intervention.url = NULL;
    intervention.log = NULL;
    intervention.disruptive = 0;
    return msc_intervention(transaction, &intervention);
}



// Helper functions to store all core rule set file names in memory.
static char **makeCharArray(int size) {
    //fprintf(stderr, "  C makeCharArray\n");
    return calloc(sizeof(char*), size);
}
static void freeCharArray(char **array, int size) {
    int index;
    //fprintf(stderr, "  C freeCharArray\n");
    for (index = 0; index < size; index++)
    {
        free(array[index]);
    }

    free(array);
}
static void setArrayString(char **array, char *filename, int index) {
    //fprintf(stderr, "setArrayString [%d] = '%s' start\n", index, file);
    array[index] = filename;
    //fprintf(stderr, "setArrayString [%d] = '%s' -end-\n", index, file);
}
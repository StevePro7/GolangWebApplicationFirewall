#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// Helper function to initialize ModSec.
static void initializeModSecurityImpl();

int LoadModSecurityCoreRuleSet(char **array, int size)
{
    int index = 0;
    const char *file;
    const char *error = NULL;
    if (modsec == NULL)
    {
        initializeModSecurityImpl();
    }

    for( index = 0; index < size; index++ )
    {
        file = array[index];
        msc_rules_add_file(rules, file, &error);
        if (error != NULL)
        {
            break;
        }
    }

    return index;
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

void InitializeModSecurity()
{
    ModSecurity *modsec = NULL;
    RulesSet *rules = NULL;

    initializeModSecurityImpl();
}
static void initializeModSecurityImpl()
{
    modsec = msc_init();
    rules = msc_create_rules_set();
}

// Helper functions to store all core rule set file names in memory.
static char **makeCharArray(int size) {
    return calloc(sizeof(char*), size);
}
static void freeCharArray(char **array, int size) {
    int index;
    for (index = 0; index < size; index++)
    {
        free(array[index]);
    }

    free(array);
}
static void setArrayString(char **array, char *filename, int index) {
    array[index] = filename;
}
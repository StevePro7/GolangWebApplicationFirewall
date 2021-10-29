#include <stdio.h>
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

void MyCInit()
{
    //printf("MyCInit start\n");
    if (modsec != NULL)
    {
        return;
    }

    modsec = msc_init();
    rules = msc_create_rules_set();

    const char *error = NULL;
    msc_rules_add_file(rules, "modsecdefault.conf", &error);
    msc_rules_add_file(rules, "crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
    if (error != NULL)
    {
        printf("msc_rules_add_file error : '%s'\n", error);
    }

    //msc_rules_dump(rules);
    //printf("MyCInit -end-\n");
    return;
}

//int MyCProcess(char *uri, char *protocol, char *version)
//int MyCProcess(char *uri)
int MyCProcess(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port)
{
    Transaction *transaction = NULL;
    transaction = msc_new_transaction(modsec, rules, NULL);

    fprintf(stderr, "URI='%s'\n", uri);
    fprintf(stderr, "Method='%s'\n", http_method);
    fprintf(stderr, "Protocol='%s/%s'\n", http_protocol, http_version);
    fprintf(stderr, "Client socket='%s:%d'\n", client_link, client_port);
    fprintf(stderr, "Server socket='%s:%d'\n", server_link, server_port);

    msc_process_connection(transaction, "127.0.0.1", 80, "127.0.0.1", 80);
    msc_process_uri(transaction, uri, http_protocol, http_version);
    msc_process_request_headers(transaction);
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

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

int MyCProcess(char *uri, char *body)
{
    Transaction *transaction = NULL;
    transaction = msc_new_transaction(modsec, rules, NULL);
    msc_process_connection(transaction, "127.0.0.1", 80, "127.0.0.1", 80);
    fprintf(stderr, "URI=%s\n", uri);
    msc_process_uri(transaction, uri, "CONNECT", "1.1");
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

/*
int MyCProcess(char *uri, char *body)
{
    printf("MyCProcess start [%s]\n", uri);
    int inter = 0;
    Transaction *transaction = NULL;
    transaction = msc_new_transaction(modsec, rules, NULL);
    msc_process_connection(transaction, "127.0.0.1", 80, "127.0.0.1", 80);
    fprintf(stderr, "URI='%s'\n", uri);

    ModSecurityIntervention intervention;
    intervention.status = 200;
    intervention.url = uri;
    intervention.log = NULL;
    intervention.disruptive = 0;

    printf("intervention URL [%s]\n", intervention.url);

    printf("... msc_intervention START\n");
    inter = msc_intervention(transaction, &intervention);
    printf("... msc_intervention -END-\n");
    fprintf(stderr, "intervention=%d\n", inter);


    msc_process_uri(transaction, uri, "GET", "1.1");
    printf("...connect\n");
    //msc_process_uri(transaction, uri, "CONNECT", "2.0");

//    printf("... msc_intervention START\n");
//    inter = msc_intervention(transaction, &intervention);
//    printf("... msc_intervention -END-\n");
//    fprintf(stderr, "intervention=%d\n", inter);


//    msc_append_request_body(transaction, body, 10);
//    printf("...msc_append_request_body ['%s']\n", body);
//    printf("...\n");
//
//    printf("... msc_intervention START\n");
//    inter = msc_intervention(transaction, &intervention);
//    printf("... msc_intervention -END-\n");
//    fprintf(stderr, "intervention=%d\n", inter);


    msc_process_request_headers(transaction);
    printf("...msc_process_request_headers\n");
    printf("...\n");

    printf("... msc_intervention START\n");
    inter = msc_intervention(transaction, &intervention);
    printf("... msc_intervention -END-\n");
    fprintf(stderr, "intervention=%d\n", inter);


    printf("...msc_process_request_body start\n");
    msc_process_request_body(transaction);
    printf("...msc_process_request_body -end-\n");
    printf("...\n");

    printf("... msc_intervention START\n");
    inter = msc_intervention(transaction, &intervention);
    printf("... msc_intervention -END-\n");
    fprintf(stderr, "intervention=%d\n", inter);


    printf("...msc_process_response_headers -end-\n");
    printf("...\n");
    msc_process_response_headers(transaction, 200, "HTTP 1.1");

    printf("...msc_process_response_body -end-\n");
    msc_process_response_body(transaction);

    printf("...msc_process_logging -end-\n");
    msc_process_logging(transaction);    //msc_process_response_headers(transaction);
    //emsc_process_response_body(transaction);



    printf("... msc_intervention START\n");
    inter = msc_intervention(transaction, &intervention);
    printf("... msc_intervention -END-\n");
    fprintf(stderr, "intervention=%d\n", inter);

//    printf("val1 %d\n", intervention.status);
//    printf("val2 %s\n", intervention.url);
//    printf("val3 %s\n", intervention.log);
//    printf("val4 %d\n", intervention.disruptive);

    printf("MyCProcess -end-!! [%s]\n", uri);
    return inter;
}
*/
package main

/*
#cgo CPPFLAGS: -I/usr/local/modsecurity/include
#cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
#include "modsecurity/rules_set.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"
ModSecurity *modsec = NULL;
RulesSet *rules = NULL;
void init() {
    if (modsec!=NULL) {
        return;
    }
    modsec = msc_init();
    rules = msc_create_rules_set();
    const char *error = NULL;
    msc_rules_add_file(rules, "modsecdefault.conf", &error);
    msc_rules_add_file(rules, "crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
    return;
}
int process(const char *uri)
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
    return inter;
}
*/

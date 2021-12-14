#include "waf.h"
#include "_cgo_export.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// Helper function to initialize ModSec.
static void initializeModSecurityImpl();


void call_sgb(texts funcs, char *str)
{
    printf("C  call_sgb beg\n");
    printf("C  call_sgb '%s'\n", str);
    funcs(str);
    printf("C  call_sgb end\n");
}

void MyFuncPtr(void *foo, const void *bar)
{
    char *sgb = (char *)bar;
    fprintf(stderr, "MyFuncPtr start\n");
    //fprintf(stderr, "MyFuncPtr '%s'\n", foo);
    fprintf(stderr, "MyFuncPtr '%s' [SPLAT..!]\n", sgb);
    call_sgb(&GoText, sgb);
    fprintf(stderr, "MyFuncPtr -end-..END..!\n");
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
    fprintf(stderr, "  C msc_set_log_cb start\n");
    msc_set_log_cb(modsec, MyFuncPtr);
    fprintf(stderr, "  C msc_set_log_cb -end-\n");
    rules = msc_create_rules_set();
}

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

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size)
{
    return calloc(sizeof(char*), size);
}
void freeCharArray(char **array, int size)
{
    int index;
    for (index = 0; index < size; index++)
    {
        free(array[index]);
    }

    free(array);
}
void setArrayString(char **array, char *filename, int index)
{
    array[index] = filename;
}


//void call_const(ModSecLogCb callback)
//{
//    //char *str = "steven";
//    //const char *buf = "World";
//    //char *buf = "boland";
//    printf("C  call_const beg\n");
//    //callback(str, buf);
//    //callback(str);
//    printf("C  call_const end!!\n");
//}
// void pass_GoConst(void)
// {
//    printf("C  pass_GoConst beg\n");
//    call_const(&GoConst);
//    printf("C  pass_GoConst end\n");
// }
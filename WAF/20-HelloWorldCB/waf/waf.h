#include <stdlib.h>

typedef long long int(*adder)(long long int, long long int);
typedef void(*texts)(char*);

extern void call_add(adder, int, int);
extern void call_sgb(texts, char*);

void InitializeModSecurity();
int LoadModSecurityCoreRuleSet(char **array, int size);
int ProcessHttpRequest(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port);

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size);
void freeCharArray(char **array, int size);
void setArrayString(char **array, char *filename, int index);


//extern void call_const(ModSecLogCb);
//extern void pass_GoConst(void);

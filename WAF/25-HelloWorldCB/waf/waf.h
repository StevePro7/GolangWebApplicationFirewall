// Important: <stdlib.h> header file required for Golang statements e.g. defer C.free(unsafe.Pointer(API))
#include <stdlib.h>

typedef void(*texts)(char*);
extern void call_sgb(texts, char*);

void InitializeModSecurity();
int LoadModSecurityCoreRuleSet(char **array, int size);
int ProcessHttpRequest(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port);

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size);
void freeCharArray(char **array, int size);
void setArrayString(char **array, char *filename, int index);

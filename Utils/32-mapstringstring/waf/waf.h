// IMPORTANT - required for defer in Go code
// defer C.free(unsafe.Pointer(API))
#include <stdlib.h>
int AddRequestHeaders(char **reqHeaderKeys, char **reqHeaderVals, int reqHeaderSize, char *reqBodyText, int reqBodySize);

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size);
void freeCharArray(char **array, int size);
void setArrayString(char **array, char *input, int index);

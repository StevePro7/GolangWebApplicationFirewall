typedef long long int(*adder)(long long int, long long int);
typedef void (*ModSecLogCb) (void *, const void *);

extern void call_add(adder);
extern void pass_GoAdd(void);

extern void call_text(ModSecLogCb);
extern void pass_GoText(void);

typedef long long int(*adder)(long long int, long long int);
typedef void (*ModSecNumCb) (int);
//typedef void (*ModSecLogCb) (void *, const void *);
//typedef void (*ModSecLogCb) (char *);
typedef void (*ModSecLogCb) (char *);
typedef void (*ModSecConstCb) (void *, void *);
//typedef void (*ModSecConstCb) (void *);

extern void call_add(adder);
extern void pass_GoAdd(void);
extern void call_num(ModSecNumCb);
extern void pass_GoNum(void);

extern void call_text(ModSecLogCb);
extern void pass_GoText(void);
extern void call_const(ModSecConstCb);
extern void pass_GoConst(void);

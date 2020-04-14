#include <sys/types.h>
#include <stropts.h>
#include <stdlib.h>
#include <memory.h>

int
senddata(int fd, char *buf, uint_t blen, char *addr,
     ushort_t alen)
{
     struct data_req *reqp;
     struct strbuf ctl, dat;
     char *bp;
     int size, ret;

#include <sys/types.h>
#include <stropts.h>
#include <unistd.h>
#include <errno.h>

int
getack(int fd)
{
     struct data_ack ack;
     struct strbuf ctl;
     int fl = RS_HIPRI;
     int ret;

#include <poll.h>
#include <unistd.h>

extern void error(const char *fmt, ...);

void
comm(int tfd, int nfd)
{
     int n, i;
     struct pollfd pfd[2];
     char buf[256];

#include <poll.h>
#include <unistd.h>

extern void error(const char *fmt, ...);

void
comm(int tfd, int nfd)
{
     int n, i;
     struct pollfd pfd[2];
     char buf[256];
     pfd[0].fd = tfd;    /* terminal */
     pfd[0].events = POLLIN;
     pfd[1].fd = nfd;    /* network */
     pfd[1].events = POLLIN;
     for (;;) {
         /*
          * Wait for events to occur.
          */
         if (poll(pfd, 2, -1) < 0) {
             error("poll failed");
             break;
         }

#include <sys/types.h>
#include <stdlib.h>
#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <errno.h>
#include <stropts.h>
#include <signal.h>

#define BUFSIZE (64*1024)
#define RDSIZE  (BUFSIZE/8)

char buf[BUFSIZE];
int widx, ridx;      /* write and read indices */
int totwr, totrd;    /* total amounts read and written */
int flowctl;         /* 1 if flow-controlled, 0 if not */
int nfc;             /* number of times flow-controlled */

void catreg(int), cattostream(int);
int doread(int);
void dowrite(int), finwrite(void);
void setblock(int), setnonblock(int);

#ifdef FCBUG
void nop(int);
#endif

extern void error(const char *fmt, ...);
extern void fatal(const char *fmt, ...)

void
main(int argc, char *argv[])
{
     int i, fd, isoutstr;
#ifdef FCBUG
     struct sigaction sa;

     /*
      * If system contains flow-control bug,
      * install a signal handler for SIGALRM.
      */
     sa.sa_handler = nop;
     sigemptyset(&sa.sa_mask);
     sa.sa_flags = 0;

     if (sigaction(SIGALRM, &sa, NULL) < 0)
         fatal("cat: sigaction failed");
#endif

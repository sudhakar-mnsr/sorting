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

     /*
      * See if the standard output is a stream.  If
      * isastream fails, assume stdout is not a stream.
      */
     isoutstr = isastream(1);
     if (isoutstr == -1)
         isoutstr = 0;

     /*
      * Process each file named on the command line.
      */
     for (i = 1; i < argc; i++) {
         if ((fd = open(argv[i], O_RDONLY)) < 0) {
             error("cat: cannot open %s", argv[i]);
             continue;
         }
         /*
          * If the standard output is a stream, call
          * cattostream to print the file.  Otherwise
          * call catreg (see Example 2.4.6) to do it.
          */
         if (isoutstr)
             cattostream(fd);
         else
             catreg(fd);
         close(fd);
     }
#ifdef DEBUG
     printf("cat: number of flow controls = %d\n", nfc);
#endif
     exit(0);
}

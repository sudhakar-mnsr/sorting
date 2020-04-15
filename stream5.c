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
extern void fatal(const char *fmt, ...);

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

#ifdef FCBUG
void
nop(int sig)
{
}
#endif

void
cattostream(int fd)
{
     int n;
     struct sigaction sa;
     sigset_t s, os;

     sigemptyset(&s);
     sigaddset(&s, SIGPOLL);

     /*
      * Install a signal handler for SIGPOLL.
      */
     sa.sa_handler = dowrite;
     sigemptyset(&sa.sa_mask);
     sa.sa_flags = 0;
     if (sigaction(SIGPOLL, &sa, NULL) < 0)
         fatal("cat: sigaction failed");
     /*
      * Arrange to be notified when the standard output
      * is no longer flow-controlled.  Then place the file
      * descriptor for stdout in nonblocking mode.
      */
     if (ioctl(1, I_SETSIG, S_OUTPUT) < 0)
         fatal("cat: I_SETSIG ioctl failed");
     setnonblock(1);
     totrd = totwr = 0;
     ridx = widx = 0;
     flowctl = 0;

     for (;;) {
         if ((n = doread(fd)) == 0) {
             /*
              * End of file; finish writing.
              */
             finwrite();
             break;
         } else if (n < 0) {
             /*
              * Read was interrupted by SIGPOLL.
              */
             continue;
         } else {
             /*
              * Successfully read something.
              */
             totrd += n;
         }
         /*
          * Start critical section.  Block SIGPOLL.

          * Then try to write what we’ve just read.
          */
         sigprocmask(SIG_BLOCK, &s, &os);
         dowrite(0);
         while (flowctl) {
             if (ridx != widx) {
                 /*
                  * Allow read to be interrupted.
                  */
                 sigprocmask(SIG_UNBLOCK, &s, NULL);
                 if ((n = doread(fd)) == 0) { /* EOF */
                     finwrite();
                     return;
                 } else if (n > 0) { /* read data */
                     totrd += n;
                 }
                 sigprocmask(SIG_BLOCK, &s, NULL);
             } else {
#ifdef FCBUG
                 /*
                  * Flow control bug -- might miss event.
                  */
                 alarm(1);
#endif

                 /*
                  * Atomically unblock SIGPOLL and
                  * wait to be interrupted.  On return,
                  * SIGPOLL is still blocked.
                  */
                 sigsuspend(&os);
#ifdef FCBUG
                 alarm(0);
                 if (ioctl(1, I_CANPUT, 0) != 0) {
                     /*
                      * Flow control lifted;
                      * continue writing.
                      */
                     flowctl = 0;
                     dowrite(0);
                     break;
                 }
#endif
             }
         }

         /*
          * End critical section.  Unblock SIGPOLL.
          */
         sigprocmask(SIG_UNBLOCK, &s, NULL);
     }
}
void
setnonblock(int fd)
{
     int fl;

     /*
      * Get the current file flags and turn on
      * nonblocking mode.
      */
     if ((fl = fcntl(fd, F_GETFL, 0)) < 0)
         fatal("cat: fcntl F_GETFL failed");
     if (fcntl(fd, F_SETFL, fl|O_NONBLOCK) < 0)
         fatal("cat: fcntl F_SETFL failed");
}
void
setblock(int fd)
{
     int fl;

     /*
      * Get the current file flags and turn off
      * nonblocking mode.
      */
     if ((fl = fcntl(fd, F_GETFL, 0)) < 0)
         fatal("cat: fcntl F_GETFL failed");
     if (fcntl(fd, F_SETFL, (fl&~O_NONBLOCK)) < 0)
         fatal("cat: fcntl F_SETFL failed");
}
void
finwrite(void)
{

     /*
      * Cancel SIGPOLL generation for stdout.
      */
     if (ioctl(1, I_SETSIG, 0) < 0) {
         struct sigaction sa;
         /*
          * I_SETSIG shouldn’t have failed, but
          * it did, so the next best thing is to
          * ignore SIGPOLL.
          */
         sa.sa_handler = SIG_IGN;
         sigemptyset(&sa.sa_mask);
         sa.sa_flags = 0;
         if (sigaction(SIGPOLL, &sa, NULL) < 0)
             fatal("sigaction failed");
     }
     /*
      * Disable nonblocking mode and write last
      * portion to the standard output.
      */
     setblock(1);
     dowrite(0);
}
int
doread(int fd)
{
     int n, rcnt;

     /*
      * Calculate the space left to read.
      * Read at most RDSIZE bytes.
      */
     rcnt = widx - ridx;
     if (rcnt <= 0) {
         /*
          * The writer is behind the reader
          * in the buffer.
          */
         rcnt = BUFSIZE - ridx;
         if (rcnt > RDSIZE)

             rcnt = RDSIZE;
     }

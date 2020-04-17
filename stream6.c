#include <sys/types.h>
#include <unistd.h>
#include <signal.h>
#include <stropts.h>

int
send(int fd, char *buf, uint_t blen, char *addr,
     uint_t alen)
{
     sigset_t set, oset;

     /*
      * Block SIGPOLL.
      */
     sigemptyset(&set);
     sigaddset(&set, SIGPOLL);
     sigprocmask(SIG_BLOCK, &set, &oset);
     /*
      * Send the message.
      */
     if (senddata(fd, buf, blen, addr, alen) < 0) {
         sigprocmask(SIG_SETMASK, &oset, NULL);
         return(-1);
     }

#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

#define TTYMODE  (S_IRUSR|S_IWUSR|S_IWGRP)
int
chgterm()
{
     int pfd[2];
     char *tty;

     /*
      * Get the name of the controlling terminal.
      */
     if ((tty = ttyname(0)) == NULL)
         return(-1);
     /*
      * Create a pipe and mount one end on top of
      * the terminal’s device node.  Then change
      * the mode of the pipe to give it the same
      * permissions as terminals.
      */
     if (pipe(pfd) < 0)
         return(-1);
     if ((fattach(pfd[1], tty) < 0) ||
       (chmod(tty, TTYMODE) < 0)) {
         close(pfd[0]);
         close(pfd[1]);
         return(-1);
     }
     /*
      * Close the end of the pipe just mounted and
      * return the other end to the caller.
      */
     close(pfd[1]);
     return(pfd[0]);
}

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

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

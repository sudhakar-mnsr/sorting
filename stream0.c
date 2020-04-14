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

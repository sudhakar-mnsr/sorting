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

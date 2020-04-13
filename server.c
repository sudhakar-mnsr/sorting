#include <stdio.h>
#include <sys/types.h>
#include <sys/fcntl.h>
#include <sys/errno.h>

#include <sys/socket.h>
#include <netinet/in.h>

extern int errno;

#define SERV_TCP_PORT   6000

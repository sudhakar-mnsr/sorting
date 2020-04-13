#include <stdio.h>
#include <sys/types.h>
#include <sys/fcntl.h>

#include <sys/socket.h>
#include <netinet/in.h>

extern int errno;

#define SERV_HOST_ADDR  "128.220.101.4"
#define SERV_TCP_PORT   6000

main()
{
	int fd;
	struct sockaddr_in	my_addr;
	struct sockaddr_in	serv_addr;
	char buf[128];
	void echo_driver(int, struct sockaddr_in *);

	if ((fd = net_open ("/dev/udp", O_RDWR)) < 0)
	{
		fprintf (stderr, "open failed.\n");
		exit (-1);
	}

/* include icmpd1 */
#include	"icmpd.h"

int
main(int argc, char **argv)
{
	int		i, sockfd;
	struct sockaddr_un sun;

	if (argc != 1)
		err_quit("usage: icmpd");

	maxi = -1;					/* index into client[] array */
	for (i = 0; i < FD_SETSIZE; i++)
		client[i].connfd = -1;	/* -1 indicates available entry */
	FD_ZERO(&allset);

	fd4 = Socket(AF_INET, SOCK_RAW, IPPROTO_ICMP);
	FD_SET(fd4, &allset);
	maxfd = fd4;

/* include icmpd1 */
#include	"icmpd.h"

int
main(int argc, char **argv)
{
	int		i, sockfd;
	struct sockaddr_un sun;

	if (argc != 1)
		err_quit("usage: icmpd");

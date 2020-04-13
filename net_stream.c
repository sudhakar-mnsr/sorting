#include <stdio.h>
#include <sys/types.h>
#include <sys/fcntl.h>
#include <sys/errno.h>

#include <sys/stream.h>
#include <sys/stropts.h>
#include <sys/tihdr.h>


extern int errno;

int
net_open (char *path, int oflags, void *addr, int addrlen)
{
	int fd;
	int flags;

	if ((fd = open (path, oflags)) < 0)
	{
		perror ("open");
		return (-1);
	}
	return (fd);
}

int
net_bind (int fd, void *addr, int addrlen)
{
	struct {
		struct T_bind_req msg_hdr;
		char addr[128];
	} bind_req;
	struct strbuf ctlbuf;
	union T_primitives rcvbuf;
	struct T_error_ack *error_ack;
	int flags;

	if (addr == NULL || addrlen == 0)
	{
		fprintf (stderr, "No address\n");
		return (-1);
	}
	bind_req.msg_hdr.PRIM_type = T_BIND_REQ;
	bind_req.msg_hdr.ADDR_length = addrlen;
	bind_req.msg_hdr.ADDR_offset = sizeof (struct T_bind_req);
	bind_req.msg_hdr.CONIND_number = 5;
	bcopy (addr, bind_req.addr, addrlen);

	ctlbuf.len = sizeof (struct T_bind_req) + addrlen;
	ctlbuf.buf = (char *) &bind_req;

	if (putmsg (fd, &ctlbuf, NULL, 0) < 0)
	{
		return (-1);
	}
	/*
	 * Wait for acknowledgement
	 */
	ctlbuf.maxlen = sizeof (union T_primitives);
	ctlbuf.len = 0;
	ctlbuf.buf = (char *) &rcvbuf;
	flags = RS_HIPRI;
	if (getmsg (fd, &ctlbuf, NULL, &flags) < 0)
	{
		perror ("getmsg");
		return (-1);
	}

	if (ctlbuf.len < sizeof (long))
	{
		fprintf (stderr, "Bad length from getmsg.\n");
		errno = EPROTO;
		return (-1);
	}

	switch (rcvbuf.type)
	{
	    case T_BIND_ACK:
		return (fd);
	    case T_ERROR_ACK:
		if (ctlbuf.len < sizeof (struct T_error_ack))
		{
			errno = EPROTO;
			return (-1);
		}
		error_ack = (struct T_error_ack *) &rcvbuf;
		fprintf (stderr, "Error ack from bind (%d %d %d)\n",
			error_ack->ERROR_prim,
			error_ack->TLI_error,
			error_ack->UNIX_error);
		errno = error_ack->UNIX_error;
		break;
	    default:
		fprintf (stderr, "No ack from bind?\n");
		errno = EPROTO;
		break;
	}
	return (-1);
}

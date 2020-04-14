#include <sys/types.h>
#include <stropts.h>
#include <stdlib.h>
#include <memory.h>

int
senddata(int fd, char *buf, uint_t blen, char *addr,
     ushort_t alen)
{
     struct data_req *reqp;
     struct strbuf ctl, dat;
     char *bp;
     int size, ret;
     /*
      * Allocate a memory buffer large enough to hold
      * the control information.
      */
     size = sizeof(struct data_req) + alen;
     if ((bp = malloc(size)) == NULL)
         return(-1);

     /*
      * Initialize the data_req structure.
      */
     reqp = (struct data_req *)bp;
     reqp->primitive = DATA_REQUEST;
     reqp->addr_len = alen;
     reqp->addr_offset = sizeof(struct data_req);
     /*
      * Copy the address to the buffer.
      */
     memcpy(bp + reqp->addr_offset, addr, alen);
     ctl.buf = bp;
     ctl.len = size;
     dat.buf = buf;
     dat.len = blen;

     /*
      * Send the message downstream, free the memory
      * allocated for the control buffer, and return.
      */
     ret = putmsg(fd, &ctl, &dat, 0);
     free(bp);
     return(ret);
}

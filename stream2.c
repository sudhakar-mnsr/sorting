#include <sys/types.h>
#include <stropts.h>
#include <unistd.h>
#include <errno.h>

int
getack(int fd)
{
     struct data_ack ack;
     struct strbuf ctl;
     int fl = RS_HIPRI;
     int ret;
     /*
      * Initialize the control buffer and retrieve the
      * acknowledgement message.
      */
     ctl.buf = (caddr_t)&ack;
     ctl.maxlen = sizeof(struct data_ack);
     ret = getmsg(fd, &ctl, NULL, &fl);
     if (ret != 0) {
         /*
          * ret shouldn't be greater than 0, but if it
          * is, then the message was improperly formed.
          */
         if (ret > 0)
             errno = EPROTO;
         return(-1);
     }
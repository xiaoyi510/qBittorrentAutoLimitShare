FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ADD ./qbit-auto-limit $WORKDIR/qbit-auto-limit
ADD ./conf/demo.app.yaml $WORKDIR/conf/app.yaml

RUN chmod +x $WORKDIR/qbit-auto-limit

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./qbit-auto-limit

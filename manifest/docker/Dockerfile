FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ADD ./dist/qbAuto_linux_amd64/qBittorrentAutoLimitShare $WORKDIR/qBittorrentAutoLimitShare
ADD ./conf/demo.app.yaml $WORKDIR/conf/app.yaml

RUN chmod +x $WORKDIR/qBittorrentAutoLimitShare

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./qBittorrentAutoLimitShare

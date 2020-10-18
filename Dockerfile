FROM alpine
MAINTAINER hr@v2rss.ga
RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ["v2rss"]
CMD ["-p","5500"]
EXPOSE [3000,5500]
COPY ./v2rss /usr/bin/v2rss
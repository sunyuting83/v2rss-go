FROM alpine
RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ["v2rss"]
CMD ["-p","5500","-c","/home/v2list"]
COPY ./v2rss /usr/bin/v2rss
COPY ./v2list /home/v2list
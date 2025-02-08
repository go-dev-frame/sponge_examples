FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add curl, used for http service checking, can be installed without it if deployed in k8s
RUN apk add curl

COPY configs/ /app/configs/
COPY store /app/store
RUN chmod +x /app/store

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./store", "-c", "configs/store.yml"]
# if you use the Configuration Center, store.yml is changed to the Configuration Center configuration.
#CMD ["./store", "-c", "configs/store.yml", "-enable-cc"]

FROM alpine:latest

RUN apk --no-cache upgrade \
    && apk --no-cache add shadow ca-certificates \
    && addgroup -S kubewatch \
    && groupmod -g 9001 kubewatch \
    && adduser -S kubewatch -G kubewatch \
    && usermod -u 9001 kubewatch \
    && apk del shadow \
    && mkdir -p /opt/kubewatch \
    && chmod 777 /var/log

WORKDIR /opt/kubewatch

COPY ./bin/tlssetup-linux-386 ./tlssetup

RUN chmod -R 550 /opt/kubewatch

ENTRYPOINT ["./tlssetup"]

FROM golang:alpine

RUN apk --no-cache --update add make

WORKDIR /go/src/bitbucket.org/budry/release-monitor
COPY . .

VOLUME /etc/release-monitor
VOLUME /var/lib/release-monitor

RUN make

CMD ["./dist/release-monitor"]
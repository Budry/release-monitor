FROM golang:alpine

RUN apk --no-cache --update add make git

WORKDIR /go/src/bitbucket.org/budry/release-monitor

VOLUME /etc/release-monitor
VOLUME /var/lib/release-monitor

COPY . .

RUN make

CMD ["./dist/release-monitor"]
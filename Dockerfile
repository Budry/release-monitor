FROM golang:alpine as builder

RUN apk --no-cache --update add make git

WORKDIR /go/src/bitbucket.org/budry/release-monitor

VOLUME /etc/release-monitor
VOLUME /var/lib/release-monitor

COPY . .

RUN make

CMD ["./dist/release-monitor"]


###
#
###

FROM alpine

WORKDIR /usr/local/bin
COPY --from=builder /go/src/bitbucket.org/budry/release-monitor/dist/release-monitor /usr/local/bin/
CMD release-monitor
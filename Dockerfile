FROM golang:1.11 as builder

RUN mkdir -p /go/src/github.com/outbound-go
WORKDIR  /go/src/github.com/outbound-go

COPY . .
RUN make .build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/outbound-go/outbound-go .
ENTRYPOINT [ "./outbound-go" ]
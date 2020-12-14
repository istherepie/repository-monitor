FROM golang:1.15.6-alpine3.12 as build

WORKDIR /go/src/request-monitor

RUN mkdir /distribution

COPY . /go/src/request-monitor

RUN go build -o /distribution/request-monitor cmd/main.go

FROM alpine:3.12

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=build /distribution/request-monitor .

ENTRYPOINT []
CMD ["/app/request-monitor", "-host", "0.0.0.0"]

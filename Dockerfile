FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/k8s-wireguard-mgr/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/k8s-wireguard-mgr
FROM scratch
USER 10001
WORKDIR /app
COPY --from=builder /go/bin/k8s-wireguard-mgr /app/k8s-wireguard-mgr
ENTRYPOINT ["/app/k8s-wireguard-mgr"]

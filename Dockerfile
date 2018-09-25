FROM alpine:latest AS alpine
RUN apk --no-cache add ca-certificates

FROM golang:alpine AS build
WORKDIR /go/src/github.com/charliekenney23/grpc-greet-service
ENV GO111MODULE=on
COPY . .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get && mkdir -p /bin \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /bin/greeter cmd/greeter/main.go

FROM scratch
WORKDIR /bin
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=build /bin/greeter ./
EXPOSE 3000
ENTRYPOINT [ "/bin/greeter" ]

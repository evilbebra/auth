# build golang binary
FROM golang:1.19-alpine AS builder

RUN go version
RUN apk add git

COPY ./ /go/src/github.com/evilbebra/auth/
WORKDIR /go/src/github.com/evilbebra/auth/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/auth-app ./cmd/auth-app/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/api-app ./cmd/api-app/main.go

# lightweight alpine container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /go/src/github.com/evilbebra/auth/.bin/auth-app .
COPY --from=0 /go/src/github.com/evilbebra/auth/.bin/api-app .

COPY --from=0 /go/src/github.com/evilbebra/auth/config/ ./config/
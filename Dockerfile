FROM golang:1.14-alpine AS build-env

COPY . $GOPATH/src/github.com/sergushka/ahoi-kino
WORKDIR $GOPATH/src/github.com/sergushka/ahoi-kino/main
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /go/bin/ahoi/main

FROM alpine:3
RUN apk add --no-cache ca-certificates git
RUN mkdir -p creds
COPY --from=build-env /go/bin/ahoi/main /go/bin/ahoi/main
ENTRYPOINT ["/go/bin/ahoi/main"]
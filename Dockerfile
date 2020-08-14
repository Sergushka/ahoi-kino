FROM golang:latest AS build-env

LABEL maintainer="Alexander Romashin <cooladress@mail.ru>"

COPY . /ahoi-kino
WORKDIR /ahoi-kino
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -v -mod=readonly -o runserver ./main" --command=./runserver
FROM golang:bullseye

ENV GO111MODULE=on

ENV PORT=":8087"
ENV CLUSTERIP="127.0.0.1"
ENV VERSION="1.0.1"

RUN mkdir /build

COPY . /build

WORKDIR /build/src

RUN go build -o webservercicd

EXPOSE 8087

ENTRYPOINT ["./webservercicd"]


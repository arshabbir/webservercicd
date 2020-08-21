FROM golang

ENV PORT=":8087"
ENV VERSION="1.0.0"

RUN mkdir /build

COPY . /build

WORKDIR /build/src

RUN go build -o webservercicd

EXPOSE 8080

ENTRYPOINT ["./webservercicd"]


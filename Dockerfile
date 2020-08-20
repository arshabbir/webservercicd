FROM golang

ENV PORT=":8080"
ENV VERSION="1.0.0"

RUN mkdir /build

COPY . /build

WORKDIR /build

RUN go build -o webservercicd

EXPOSE 8080

ENTRYPOINT ["./webservercicd"]

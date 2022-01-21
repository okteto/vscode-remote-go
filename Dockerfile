FROM golang:buster

WORKDIR /usr/src/pp
ADD . .
RUN go build -o /usr/local/bin/hello-world

EXPOSE 8080
CMD ["/usr/local/bin/hello-world"]

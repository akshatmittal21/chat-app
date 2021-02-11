FROM golang:latest
RUN mkdir -p /go/src/chat
WORKDIR /go/src/chat
COPY server /go/src/chat
RUN go get
RUN go build -o main .
EXPOSE 3600
CMD ["/go/src/chat/main"]

FROM golang:alpine AS builder
WORKDIR /go/src/chat
COPY server .
RUN go build -o chat-server .

FROM alpine
WORKDIR /chat
COPY --from=builder /go/src/chat/chat-server /chat/ 
EXPOSE 3600
CMD ["./chat-server"]

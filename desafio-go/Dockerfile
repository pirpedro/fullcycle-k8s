FROM golang:1.14.4-alpine3.12 as builder
COPY src/hello/main.go .
RUN go build -o hello main.go

FROM alpine
COPY --from=builder /go/hello /hello
EXPOSE 8000
CMD ["/hello"]
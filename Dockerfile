FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o microserv .
FROM alpine
COPY --from=builder /build/microserv /app
WORKDIR /app
CMD ["./microserv"]